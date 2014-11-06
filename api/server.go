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
func HandleServer(w http.ResponseWriter, req *http.Request) {
	// Query server
	stat, err := server.Query.Full()

	// Set status
	status := 200
	if err != nil {
		status = 500
	}

	// Write JSON
	json.WriteJson(json.JsonWriter{
		Status: status,
		Rw:     w,
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