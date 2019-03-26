package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
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

func main() {
	port := ":1234"
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)
	}
}
