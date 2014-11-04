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
		Body:   "{}",
	})
}