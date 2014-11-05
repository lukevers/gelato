# gelato

A vanilla Minecraft HTTP JSON server.

## Building

#### 0. Before You Build

Make sure you have [Go](http://golang.org/) installed, and make sure you have set your [$GOPATH](http://golang.org/doc/code.html#GOPATH) properly, or it will prove difficult to build.

#### 1. Get The Code

Start by cloning the repository and getting the dependencies.

```bash
git clone https://github.com/lukevers/gelato
go get
```

#### 2. Build

```bash
go build
```

## Flags

#### `--conf [path/to/gelato.json]`

Specify your config file. Default is `/etc/gelato.json`

#### `--genconf`

Generate a new config file. If you don't pass the `--conf [path/to/gelato.json` flag it will by default try to create a config at `/etc/gelato.json` and will fail depending on your permissions.

#### `--port [port]`

Specify the webserver port that you want gelato to listen on. The default port is `7033`.

#### `--host [host]`

Specify the webserver host that you want gelato to listen on. The default host is `0.0.0.0`.
