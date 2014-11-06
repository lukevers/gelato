# gelato

A vanilla Minecraft HTTP JSON server.

# Building

### 0. Before You Build

Make sure you have [Go](http://golang.org/) installed, and make sure you have set your [$GOPATH](http://golang.org/doc/code.html#GOPATH) properly, or it will prove difficult to build.

### 1. Get The Code

Start by cloning the repository and getting the dependencies.

```bash
git clone https://github.com/lukevers/gelato
go get
```

### 2. Build

```bash
go build
```

# Flags

#### `--conf [path/to/gelato.json]`

Specify your config file. Default is `/etc/gelato.json`

#### `--genconf`

Generate a new config file. If you don't pass the `--conf [path/to/gelato.json` flag it will by default try to create a config at `/etc/gelato.json` and will fail depending on your permissions.

#### `--port [port]`

Specify the webserver port that you want gelato to listen on. The default port is `7033`.

#### `--host [host]`

Specify the webserver host that you want gelato to listen on. The default host is `0.0.0.0`.

# API Points

By default, gelato uses the same format for each API point. Every response will look similar to this, but with a different `Body`. All responses are valid JSON, and the comments in these blocks are for informational purposes in this README only.

```json
{
	"Status": 200, // HTTP Status Code
	"Error": null, // A string if not null
	"Body": {
		// Body
	}
}
```

## Table of Contents

* [/server](#server)
  * [/server/game-type](#server-game-type)
  * [/server/game-id](#server-game-id)
  * [/server/version](#server-version)
  * [/server/map](#server-map)
  * [/server/max-players](#server-max-players)
  * [/server/num-players](#server-num-players)
* [/players](#players)
  * [/players/online](#players-online)

### /server

The API point `/server` contains general information about the server. Information about the max number of players the server allows, and the number of current online players are included in this API point for convenience, but the `/players` API point is player specific and includes a list of current online players. The body of the response could be similar to this:

```json
"Body": {
	"GameType": "SMP",
	"GameId": "MINECRAFT",
	"Version": "1.8",
	"Map": "world",
	"MaxPlayers": "20",
	"NumPlayers": "0",
	"Motd": "A Minecraft Server"
}
```

### /players

The API point `/players` contains information about current online players. It also contains information about the max number of players the server allows. The body of the response could be similar to this:

```json
"Body": {
	"NumPlayers": 1,
	"MaxPlayers": 20,
	"OnlinePlayers": ["player1"] // If no players, this is null
}
```

### /players/online

The API point `/players/online` contains a string array of current online players. If there are no players online the result is null. The body of the response could be similar to this:

```json
"Body": {
	"OnlinePlayers": ["player1"] // If no players, this is null
}
```
