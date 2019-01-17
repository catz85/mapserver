package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"mapserver/coords"
	"time"
	"strings"
	"errors"
)

const migrateScript = `
alter table blocks add mtime integer default 0;
create index blocks_mtime on blocks(mtime);

CREATE TRIGGER update_blocks_mtime_insert after insert on blocks for each row
begin
update blocks set mtime = strftime('%s', 'now') where pos = new.pos;
end;

CREATE TRIGGER update_blocks_mtime_update after update on blocks for each row
begin
update blocks set mtime = strftime('%s', 'now') where pos = old.pos;
end;
`

//TODO: initial run: https://stackoverflow.com/questions/14468586/efficient-paging-in-sqlite-with-millions-of-records
//TODO: postgres test

type Sqlite3Accessor struct {
	db       *sql.DB
	filename string
}

func (db *Sqlite3Accessor) Migrate() error {

	//RW connection
	rwdb, err := sql.Open("sqlite3", db.filename+"?mode=rw")
	if err != nil {
		return err
	}

	defer rwdb.Close()

	hasMtime := true
	_, err = rwdb.Query("select max(mtime) from blocks")
	if err != nil {
		hasMtime = false
	}

	if !hasMtime {
		log.WithFields(logrus.Fields{"filename": db.filename}).Info("Migrating database")
		start := time.Now()
		_, err = rwdb.Exec(migrateScript)
		if err != nil {
			return err
		}
		t := time.Now()
		elapsed := t.Sub(start)
		log.WithFields(logrus.Fields{"elapsed": elapsed}).Info("Migration completed")
	}

	return nil
}

func convertRows(pos int64, data []byte, mtime int64) Block {
	c := coords.PlainToCoord(pos)
	return Block{Pos: c, Data: data, Mtime: mtime}
}

func (db *Sqlite3Accessor) FindLatestBlocks(mintime int64, limit int) ([]Block, error) {
	return make([]Block, 0), nil
}

const getBlockQuery = `
select pos,data,mtime from blocks b where b.pos = ?
`

func (db *Sqlite3Accessor) GetBlock(pos coords.MapBlockCoords) (*Block, error) {
	ppos := coords.CoordToPlain(pos)

	rows, err := db.db.Query(getBlockQuery, ppos)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var pos int64
		var data []byte
		var mtime int64

		err = rows.Scan(&pos, &data, &mtime)
		if err != nil {
			return nil, err
		}

		mb := convertRows(pos, data, mtime)
		return &mb, nil
	}

	return nil, nil
}


func (db *Sqlite3Accessor) CountBlocks(pos1 coords.MapBlockCoords, pos2 coords.MapBlockCoords) (int, error){

	poslist := make([]interface{}, 0)

	if pos1.X != pos2.X {
		return 0, errors.New("x does not line up")
	}

	if pos1.Z != pos2.Z {
		return 0, errors.New("z does not line up")
	}

	minY := pos1.Y
	maxY := pos2.Y

	if minY > maxY {
		minY, maxY = maxY, minY
	}

	for y := minY; y <= maxY; y++ {
		poslist = append(poslist, coords.CoordToPlain(coords.NewMapBlockCoords(pos1.X, y, pos1.Z)))
	}

	getBlocksQuery := "select count(*) from blocks b where b.pos in (?" + strings.Repeat(",?", len(poslist)-1) + ")"

	rows, err := db.db.Query(getBlocksQuery, poslist...)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	if rows.Next() {
		var count int

		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}

		return count, nil
	}

	return 0, errors.New("No rows returned")
}

func NewSqliteAccessor(filename string) (*Sqlite3Accessor, error) {
	db, err := sql.Open("sqlite3", filename+"?mode=ro")
	if err != nil {
		return nil, err
	}

	sq := &Sqlite3Accessor{db: db, filename: filename}
	return sq, nil
}
