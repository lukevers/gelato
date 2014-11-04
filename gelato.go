package main

import (
	"errors"
	"flag"
	"github.com/gorilla/mux"
	"github.com/lukevers/gelato/json"
	"github.com/lukevers/golem"
	"net/http"
	"strconv"
)

func main() {
	// Parse flags
	flag.Parse()

	// Create HTTP server
	r := mux.NewRouter()

	// Setup routes
	r.HandleFunc("/", HandleRoot)

	// Handle Other (all 404s)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.WriteJson(json.JsonWriter{
			Status: 404,
			Rw:     w,
			Error:  errors.New("Not Found"),
		})
	})

	// Start HTTP server
	golem.Infof("Starting HTTP server on %s", *host+":"+strconv.Itoa(*port))
	http.Handle("/", r)
	http.ListenAndServe(*host+":"+strconv.Itoa(*port), nil)
}
