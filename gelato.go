package main

import (
	"errors"
	"flag"
	"github.com/gorilla/mux"
	"github.com/lukevers/gelato/api"
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
	server = minecraft.Create(config.Host, config.RconPass, config.RconPort, config.QueryPort)

	// Pass our server instance to our API
	api.SetServer(server)

	// Create HTTP server
	r := mux.NewRouter()

	// Players API Routes
	r.HandleFunc("/players", api.HandlePlayers).Methods("GET")
	r.HandleFunc("/players/online", api.HandleOnlinePlayers).Methods("GET")

	// Server API Routes
	r.HandleFunc("/server", api.HandleServer).Methods("GET")

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
