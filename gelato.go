package main

import (
	"errors"
	"flag"
	"github.com/gorilla/mux"
	"github.com/lukevers/gelato/config"
	"github.com/lukevers/gelato/json"
	"github.com/lukevers/gelato/minecraft"
	"github.com/lukevers/golem"
	"net/http"
	"os"
	"strconv"
)

var server *minecraft.Server

func main() {
	// Parse flags
	flag.Parse()

	// Check if we want to generate a config
	if *genc {
		// Use the path from our "conf" flag
		err := config.Gen(*conf)
		if err != nil {
			golem.Warnf("Could not generate config file: %s", err)
			os.Exit(1)
		}
	}

	// Load configuration file
	err, config := config.Load(*conf)
	if err != nil {
		golem.Warnf("Could not load config file: %s", err)
		os.Exit(1)
	}

	// Create a Minecraft server
	server = minecraft.Create(config.RconHost, config.RconPort, config.RconPass)

	// Create HTTP server
	r := mux.NewRouter()

	// Setup routes
	r.HandleFunc("/", HandleRoot)
	r.HandleFunc("/players", HandlePlayers)

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
