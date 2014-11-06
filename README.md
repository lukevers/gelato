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

If you don't care about `Status` and `Error` and you only want the contents of `Body` there's some good news! By appending the query `body=true` to your request you can do that!

```json
{
	// Body
}
```

## Table of Contents

* [/server](#server)
  * [/server/gametype](#servergametype)
  * [/server/gameid](#servergameid)
  * [/server/version](#serverversion)
  * [/server/map](#servermap)
  * [/server/maxplayers](#servermaxplayers)
  * [/server/numplayers](#servernumplayers)
  * [/server/motd](#servermotd)
* [/players](#players)
  * [/players/online](#playersonline)

### /server

The API point `/server` contains general information about the server. Information about the max number of players the server allows, and the number of current online players are included in this API point for convenience, but the `/players` API point is player specific and includes a list of current online players. The body of the response could be similar to this:

```json
{
	"GameType": "SMP",
	"GameId": "MINECRAFT",
	"Version": "1.8",
	"Map": "world",
	"MaxPlayers": "20",
	"NumPlayers": "0",
	"Motd": "A Minecraft Server"
}
```

### /server/gametype

The API point `/server/gametype` contains the current [gametype]() of the server. The body of the response could be similar to this:

```json
{
	"GameType": "SMP"
}
```

### /server/gameid

The API point `/server/gameid` contains the current [gameid]() of the server. The body of the response could be similar to this:

```json
{
	"GameId": "MINECRAFT"
}
```

### /server/version

The API point `/server/version` contains the version of the Minecraft server that is running. The body of the response could be similar to this:

```json
{
	"Version": "1.8"
}
```

### /server/map

The API point `/server/map` contains the name of the Minecraft server map. The body of the response could be similar to this:

```json
{
	"Map": "world"
}
```

### /server/maxplayers

The API point `/server/maxplayers` contains the max number of players that the Minecraft server allows. The body of the response could be similar to this:

```json
{
	"MaxPlayers": 20
}
```

### /server/numplayers

The API point `/server/numplayers` contains the number of currently online players. The body of the response could be siliar to this:

```json
{
	"NumPlayers": 0
}
```

### /server/motd

The API point `/server/motd` contains the [motd]() for the Minecraft server. The body of the response could be similar to this:

```json
{
	"Motd": "A Minecraft Server"
}
```

### /players

The API point `/players` contains information about current online players. It also contains information about the max number of players the server allows. The body of the response could be similar to this:

```json
{
	"NumPlayers": 1,
	"MaxPlayers": 20,
	"OnlinePlayers": ["player1"] // If no players, this is null
}
```

### /players/online

The API point `/players/online` contains a string array of current online players. If there are no players online the result is null. The body of the response could be similar to this:

```json
{
	"OnlinePlayers": ["player1"] // If no players, this is null
}
```
