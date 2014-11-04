package main

import (
	"flag"
)

// Webserver flags
var port = flag.Int("port", 7033, "Port for webserver to bind to")
var host = flag.String("host", "0.0.0.0", "Host for webserver to bind to")
