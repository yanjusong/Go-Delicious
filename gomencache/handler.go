package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

const (
	SET byte = 1
	GET byte = 2
	DEL byte = 3
	ERR byte = 50
)

func parse(buffer []byte) (op int, key string, value []byte, err error) {
	dataBytes := len(buffer)
	if dataBytes < 4 {
		errMsg := fmt.Sprintf("Illegal buffer size, size=%d", dataBytes)
		return -1, "", nil, errors.New(errMsg)
	}

	totalBytes := binary.LittleEndian.Uint32(buffer)
	if int(totalBytes) != dataBytes {
		errMsg := fmt.Sprintf("Incomplete buffer data, experted size:%d, actual size:%d", totalBytes, dataBytes)
		return -1, "", nil, errors.New(errMsg)
	}

	opflgs := int(buffer[5])
	if opflgs != 1 || opflgs != 2 {
		return -1, "", nil, errors.New("Illegal operation")
	}

	return -1, "", nil, nil
}

func handleConnection(conn net.Conn, cache *Cache) {
	connDesc := conn.RemoteAddr().String()
	fmt.Printf("Connected: %s\n", connDesc)

	defer conn.Close()

	buffer := []byte{}
	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		if err != nil {
			fmt.Println("END-> ", err)
			return
		}

		fmt.Printf("Recv %d bytes from [%s]\n", n, connDesc)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		buffer = append(buffer, readBuf[:n]...)
		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		fmt.Printf("Recv progress %d/%d of %s\n", curBytes, totalBytes, connDesc)

		if curBytes < int(totalBytes) {
			continue
		}

		// availBuffer := buffer[:totalBytes]
		// do work using `availBuffer`

		buffer = buffer[totalBytes:]
	}
}
