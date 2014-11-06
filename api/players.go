package api

import (
	"github.com/lukevers/gelato/json"
	"net/http"
)

// -----
// HandlePlayers
//
// `/players`
//
//   - Max Players
//   - Online Players
//   - Number of Online Players
//
// -----
func HandlePlayers(w http.ResponseWriter, r *http.Request) {
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
			NumPlayers    int
			MaxPlayers    int
			OnlinePlayers []string
		}{
			stat.NumPlayers,
			stat.MaxPlayers,
			stat.Players,
		},
	})
}

// -----
// HandleOnlinePlayers
//
// `/players/online`
//
//   - Online Players
//
// -----
func HandleOnlinePlayers(w http.ResponseWriter, r *http.Request) {
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
			OnlinePlayers []string
		}{
			stat.Players,
		},
	})
}
