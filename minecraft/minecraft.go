package minecraft

import (
	"github.com/lukevers/mcgoquery"
	"github.com/lukevers/mcgorcon"
)

type Server struct {
	Host      string
	RconPass  string
	RconPort  int
	QueryPort int
	Query     *mcgoquery.Client
	Rcon      mcgorcon.Client
}

// Create returns a Server and tries to connect
func Create(host, pass string, RconPort, QueryPort int) (s *Server) {
	s = &Server{
		Host:      host,
		RconPass:  pass,
		RconPort:  RconPort,
		QueryPort: QueryPort,
	}

	go s.connect()
	return
}

// Connect tries to connect to the Minecraft server.
// Instead of returning an error if the server is
// down, we try every 15 seconds to connect until
// we have successfully connected.
func (s *Server) connect() {
	var err error

	// Try to connect to the Minecraft server Rcon
	s.Rcon, err = mcgorcon.Dial(s.Host, s.RconPort, s.RconPass)
	if err != nil {
		// Try reconnecting in 15 seconds
	}

	// Try to connect to the Minecraft server Query
	s.Query, err = mcgoquery.Create(s.Host, s.QueryPort)
	if err != nil {
		// Try reconnecting in 15 seconds
	}
}

// Raw sends a raw command to the Minecraft server.
func (s *Server) Raw(command string) (string, error) {
	return s.Rcon.SendCommand(command)
}
