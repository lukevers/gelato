package main

import (
	"github.com/lukevers/gelato/json"
	"net/http"
)

// -----
// HandlePlayers
//
//   - Max Players
//   - Online Players
//   - Number of Online Players
//
// -----
func HandlePlayers(w http.ResponseWriter, req *http.Request) {
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
		Body:   struct{
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
