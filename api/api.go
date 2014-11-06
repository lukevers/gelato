package api

import (
	"github.com/lukevers/gelato/minecraft"
)

var server *minecraft.Server

// Set the minecraft server here for API usage
func SetServer(s *minecraft.Server) {
	server = s
}
