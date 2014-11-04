package main

import (
	"flag"
	"github.com/gorilla/mux"
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

	// Start HTTP server
	http.Handle("/", r)
	http.ListenAndServe(*host+":"+strconv.Itoa(*port), nil)
}
