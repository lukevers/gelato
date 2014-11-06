package api

import (
	"github.com/lukevers/gelato/json"
	"net/http"
)

// -----
// HandleServerInfo
//
//  `/server`
//
//   - GameType
//   - GameId
//   - Version
//   - Map
//   - Max Players
//   - Number of Online Players
//   - Motd
//
// -----
func HandleServer(w http.ResponseWriter, r *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Rw:     Wr(w, r),
		Status: status,
		Error:  err,
		Body: struct {
			GameType   string
			GameId     string
			Version    string
			Map        string
			MaxPlayers int
			NumPlayers int
			Motd       string
		}{
			stat.GameType,
			stat.GameID,
			stat.Version,
			stat.Map,
			stat.MaxPlayers,
			stat.NumPlayers,
			stat.MOTD,
		},
	})
}

// -----
// HandleServerGametype
//
//  `/server/gametype`
//
//   - GameType
//
// -----
func HandleServerGametype(w http.ResponseWriter, r *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Rw:     Wr(w, r),
		Status: status,
		Error:  err,
		Body: struct {
			GameType string
		}{
			stat.GameType,
		},
	})
}

// -----
// HandleServerGameid
//
//  `/server/gameid`
//
//   - GameId
//
// -----
func HandleServerGameid(w http.ResponseWriter, r *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Rw:     Wr(w, r),
		Status: status,
		Error:  err,
		Body: struct {
			GameId string
		}{
			stat.GameID,
		},
	})
}

// -----
// HandleServerVersion
//
//  `/server/version`
//
//   - Version
//
// -----
func HandleServerVersion(w http.ResponseWriter, r *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Rw:     Wr(w, r),
		Status: status,
		Error:  err,
		Body: struct {
			Version string
		}{
			stat.Version,
		},
	})
}

// -----
// HandleServerMap
//
//  `/server/map`
//
//   - Map
//
// -----
func HandleServerMap(w http.ResponseWriter, r *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Rw:     Wr(w, r),
		Status: status,
		Error:  err,
		Body: struct {
			Map string
		}{
			stat.Map,
		},
	})
}

// -----
// HandleServerMaxPlayers
//
//  `/server/maxplayers`
//
//   - MaxPlayers
//
// -----
func HandleServerMaxPlayers(w http.ResponseWriter, r *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Rw:     Wr(w, r),
		Status: status,
		Error:  err,
		Body: struct {
			MaxPlayers int
		}{
			stat.MaxPlayers,
		},
	})
}

// -----
// HandleServerNumPlayers
//
//  `/server/numplayers`
//
//   - NumPlayers
//
// -----
func HandleServerNumPlayers(w http.ResponseWriter, r *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Rw:     Wr(w, r),
		Status: status,
		Error:  err,
		Body: struct {
			NumPlayers int
		}{
			stat.NumPlayers,
		},
	})
}

// -----
// HandleServerMotd
//
//  `/server/motd`
//
//   - Motd
//
// -----
func HandleServerMotd(w http.ResponseWriter, r *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Rw:     Wr(w, r),
		Status: status,
		Error:  err,
		Body: struct {
			Motd string
		}{
			stat.MOTD,
		},
	})
}
