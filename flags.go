package main

import (
	"flag"
)

// Configuration flags
var conf = flag.String("conf", "/etc/gelato.json", "Configuration file path")
var genc = flag.Bool("genconf", false, "Generate a new configuration file.")

// Webserver flags
var port = flag.Int("port", 7033, "Port for webserver to bind to")
var host = flag.String("host", "0.0.0.0", "Host for webserver to bind to")
