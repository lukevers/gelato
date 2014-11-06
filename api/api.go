package api

import (
	"github.com/lukevers/gelato/json"
	"github.com/lukevers/gelato/minecraft"
	"net/http"
)

var server *minecraft.Server

// Set the minecraft server here for API usage
func SetServer(s *minecraft.Server) {
	server = s
}

func Wr(w http.ResponseWriter, r *http.Request) json.Wr {
	return json.Wr{
		W: w,
		R: r,
	}
}
