
STATIC_VFS=vfs/static.go
OUT_DIR=output
ENV=GO111MODULE=on

GO_LDFLAGS=-ldflags "-linkmode external -extldflags -static"

all: build-all

$(OUT_DIR):
	mkdir $@

fmt:
	go fmt ./...

test: $(OUT_DIR)
	go generate
	go build
	$(ENV) go test ./...

clean:
	rm -rf $(STATIC_VFS)
	rm -rf $(OUT_DIR)

clean-all: clean
	rm -rf mapserver.sqlite*
	rm -rf mapserver.tiles
	rm -rf mapserver.json

build-docker:
	sudo docker run --rm -it -v $(shell pwd)/:/app -v mapserver-volume:/root/go -w /app mapserver-builder make build

build-all-docker:
	sudo docker run --rm -it -v $(shell pwd)/:/app -v mapserver-volume:/root/go -w /app mapserver-builder make build-all

build: $(OUT_DIR)
	go generate
	# native
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc go build $(GO_LDFLAGS) -o $(OUT_DIR)/mapserver-linux-x86_64

build-all: $(OUT_DIR)
	go generate
	# native
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-gnu-gcc go build $(GO_LDFLAGS) -o $(OUT_DIR)/mapserver-linux-x86_64
	# apt install gcc-8-i686-linux-gnu
	CGO_ENABLED=1 GOOS=linux GOARCH=386 CC=i686-linux-gnu-gcc-7 go build $(GO_LDFLAGS) -o $(OUT_DIR)/mapserver-linux-x86
	# apt install gcc-mingw-w64
	GOARCH=386 GOOS=windows CC=i686-w64-mingw32-gcc CGO_ENABLED=1 go build -o $(OUT_DIR)/mapserver-windows-x86.exe
	GOARCH=amd64 GOOS=windows CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 go build -o $(OUT_DIR)/mapserver-windows-x86-64.exe
	# apt install gcc-5-arm-linux-gnueabihf
	GOARCH=arm GOARM=7 CC=arm-linux-gnueabihf-gcc-5 CGO_ENABLED=1 go build $(GO_LDFLAGS) -o $(OUT_DIR)/mapserver-linux-arm

#go test -cpuprofile cpu.prof -memprofile mem.prof -bench . ./mapblockrenderer/
#go tool pprof cpu.prof
#top
