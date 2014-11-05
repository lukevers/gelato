package minecraft

import (
	"github.com/lukevers/mcgorcon"
)

type Server struct {
	RconHost string
	RconPort int
	RconPass string
	mc       mcgorcon.Client
}

// Create returns a Server and tries to connect
func Create(host string, port int, pass string) (s *Server) {
	s = &Server{
		RconHost: host,
		RconPort: port,
		RconPass: pass,
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

	// Try to connect to the Minecraft server
	s.mc, err = mcgorcon.Dial(s.RconHost, s.RconPort, s.RconPass)
	if err != nil {
		// Try reconnecting in 15 seconds
	}
}

// Raw sends a raw command to the Minecraft server.
func (s *Server) Raw(command string) (string, error) {
	return s.mc.SendCommand(command)
}
