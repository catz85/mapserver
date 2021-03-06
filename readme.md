Minetest mapserver
=======

Realtime mapserver for [Minetest](https://minetest.net)

# Documentation

* [Installation](doc/install.md)
* [Mapserver mod](doc/mod.md)
* [Parameters](doc/params.md)
* [Configuration](doc/config.md)
* [Contribution](doc/contrib.md)
* [Development](doc/dev.md)
* [License](doc/license.md)
* [Changelog](doc/changelog.md)

# How it works

See: [Incremental rendering](doc/incrementalrendering.md)

# Compatibility

* Minetest 0.4.15 - 0.4.17.1
* Minetest 5.0

# Features

## Current features

* Click-and-run installation
* Initial and incremental map rendering
* Realtime rendering and map-updating
* Realtime player and world stats
* Configurable layers (default: "Base" from y -16 to 160)
* POI [markers](doc/mapobjects.md) / [mod](doc/mod.md) integration
* Protector display
* LCD Displays as markers
* Monitoring with [Prometheus](doc/prometheus.md)

## Planned Features

* Player infos (HP, breath, etc)
* Street names / Train stuff
* City borders

# Supported map-databases
The connection is auto-detected from your `world.mt`:

* Sqlite3
* PostgreSql

# Screenshots

## Web interface
<img src="./pics/web.png">

## Terminal
<img src="./pics/terminal.png">

## Map objects (as markers)
Enable/Disable those in the [Configuration](doc/config.md)
See:  [mapobjects](doc/mapobjects.md)


# Bugs

There will be bugs, please file them in the [issues](./issues) page.
