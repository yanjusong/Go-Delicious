package main

import (
	"bufio"
	"fmt"
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

	fmt.Println(cache)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go handleConnection(conn, cache)
	}
}

func handleConnection(conn net.Conn, cache *Cache) {
	connDesc := conn.RemoteAddr().String()
	fmt.Printf("Connected: %s\n", connDesc)

	defer conn.Close()

	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Client[%s]:%s\n", connDesc, data)
		conn.Write([]byte(data))
	}
}
