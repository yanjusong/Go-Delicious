package main

import (
	"net"
)

// Server represents a TCP server instance accepting connection.
// Read client's request, whitch will be parsed into readable string, then operate memory cache.
// The operation including set, get and delete.
type Server struct {
	host string
	port string
}

// Run indicates the server is doing work.
func (server *Server) Run(maxBufferSize int) error {
	listener, err := net.Listen("tcp4", server.host+":"+server.port)
	if err != nil {
		return err
	}
	defer listener.Close()

	cache := GetCache(maxBufferSize)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go handleConnection(conn, cache)
	}
}
