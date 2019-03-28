package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

// Operation type for the memory.
const (
	SET int = 1
	GET int = 2
	DEL int = 3
	ERR int = -1
)

func parse(buffer []byte) (op int, key string, value []byte, err error) {
	dataBytes := len(buffer)
	if dataBytes < 6 {
		errMsg := fmt.Sprintf("Illegal buffer size, size=%d", dataBytes)
		return -1, "", nil, errors.New(errMsg)
	}

	totalBytes := binary.LittleEndian.Uint32(buffer)
	if int(totalBytes) != dataBytes {
		errMsg := fmt.Sprintf("Incomplete buffer data, experted size:%d, actual size:%d", totalBytes, dataBytes)
		return -1, "", nil, errors.New(errMsg)
	}

	opFlags := int(buffer[4])
	if opFlags != SET && opFlags != GET && opFlags != DEL {
		return -1, "", nil, errors.New("Illegal operation")
	}
	op = int(buffer[4])

	// buffer[5] represents status, 1 means ok, only used in transfering data to client.

	// get key.
	buffer = buffer[6:]

	if len(buffer) < 4 {
		return -1, "", nil, errors.New("Illegal key")
	}
	keyBytes := binary.LittleEndian.Uint32(buffer)
	if int(keyBytes) > len(buffer) {
		return -1, "", nil, errors.New("Illegal key data")
	}

	keyArray := buffer[4:keyBytes]
	key = string(keyArray)

	// get value also.
	if opFlags == SET {
		buffer = buffer[keyBytes:]
		if len(buffer) < 4 {
			return -1, "", nil, errors.New("Illegal value")
		}

		valueBytes := binary.LittleEndian.Uint32(buffer)
		if int(valueBytes) > len(buffer) {
			return -1, "", nil, errors.New("Illegal value data")
		}
		value = append(value, buffer[4:]...)
	}

	return op, key, value, nil
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

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		fmt.Printf("Recv progress %d/%d of %s\n", curBytes, totalBytes, connDesc)

		if curBytes < int(totalBytes) {
			continue
		}

		availBuffer := buffer[:totalBytes]
		op, key, value, err := parse(availBuffer)
		if op == -1 || err != nil {
			fmt.Println("Got invalid data from client, this connection will be aborted.")
			fmt.Println(err)
			break
		}

		getBuf := []byte{}
		stat := true
		switch op {
		case SET:
			setErr := cache.Set(key, value)
			stat = (setErr == nil)
		case GET:
			getBuf, stat = cache.Get(key)
		case DEL:
			cache.Delete(key)
		default:
			stat = false
		}
		writeToClient(conn, stat, op, getBuf)

		// for next operation.
		buffer = buffer[totalBytes:]
	}
}

func writeToClient(conn net.Conn, stat bool, op int, value []byte) {
	finalBuffer := []byte{}
	totalBytes := 6
	totalBytesArray := make([]byte, 4)
	status := 0
	if stat == true {
		status = 1
	}

	if op == GET {
		totalBytes += 4
		totalBytes += len(value)
	}

	intToSlice(totalBytesArray, totalBytes)

	finalBuffer = append(finalBuffer, totalBytesArray...)
	finalBuffer = append(finalBuffer, byte(op))
	finalBuffer = append(finalBuffer, byte(status))

	if op == GET {
		valueBytesArray := make([]byte, 4)
		intToSlice(valueBytesArray, 4+len(value))
		finalBuffer = append(finalBuffer, valueBytesArray...)
		finalBuffer = append(finalBuffer, value...)
	}

	conn.Write(finalBuffer)
}

func intToSlice(buffer []byte, n int) bool {
	if len(buffer) < 4 {
		return false
	}

	buffer[0] = byte(n & 0x000000ff)
	buffer[1] = byte((n & 0x0000ff00) >> 8)
	buffer[2] = byte((n & 0x00ff0000) >> 16)
	buffer[3] = byte((n & 0xff000000) >> 24)

	return true
}
