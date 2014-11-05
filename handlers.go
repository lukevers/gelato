package main

import (
	"github.com/lukevers/gelato/json"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, req *http.Request) {
	json.WriteJson(json.JsonWriter{
		Status: 200,
		Rw:     w,
		Error:  nil,
		Body: struct {
			Test  string
			Test2 string
		}{
			"wee",
			"weeeee!",
		},
	})
}

// Handles "/players"
func HandlePlayers(w http.ResponseWriter, req *http.Request) {
	response, _ := server.Raw("list")

	w.Write([]byte(response))
}
