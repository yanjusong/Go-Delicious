package main

import (
	"bufio"
	"encoding/binary"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func makeSetBytes(key string, value []byte) []byte {
	result := []byte{}
	keyArray := []byte(key)
	totalBytes := 6 + 4 + len(keyArray) + 4 + len(value)
	totalBytesArray := make([]byte, 4)
	keyBytesArray := make([]byte, 4)
	valueBytesArray := make([]byte, 4)

	intToSlice(totalBytesArray, totalBytes)
	intToSlice(keyBytesArray, 4+len(keyArray))
	intToSlice(valueBytesArray, 4+len(value))

	result = append(result, totalBytesArray...)
	result = append(result, byte(SET), byte(0))
	result = append(result, keyBytesArray...)
	result = append(result, keyArray...)
	result = append(result, valueBytesArray...)
	result = append(result, value...)

	return result
}

func makeBytesWithType(key string, op int) []byte {
	result := []byte{}
	keyArray := []byte(key)
	totalBytes := 6 + 4 + len(keyArray)
	totalBytesArray := make([]byte, 4)
	keyBytesArray := make([]byte, 4)

	intToSlice(totalBytesArray, totalBytes)
	intToSlice(keyBytesArray, 4+len(keyArray))

	result = append(result, totalBytesArray...)
	result = append(result, byte(op), byte(0))
	result = append(result, keyBytesArray...)
	result = append(result, keyArray...)

	return result
}

func makeGetBytes(key string) []byte {
	return makeBytesWithType(key, GET)
}

func makeDeleteBytes(key string) []byte {
	return makeBytesWithType(key, DEL)
}

func TestClientSet(t *testing.T) {
	go func() {
		server := Server{"127.0.0.1", "1234"}
		server.Run(4 * 1024 * 1024)
	}()

	<-time.After(2 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	assert.Equal(t, err, nil)

	setBuffer := makeSetBytes("name", []byte("jerry"))
	wn, wok := conn.Write(setBuffer)
	assert.Equal(t, wok, nil)
	assert.Equal(t, wn, len(setBuffer))

	buffer := []byte{}
	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, err, nil)

		t.Logf("Recv %d bytes from server", n)

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		t.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

		if curBytes < int(totalBytes) {
			continue
		}

		assert.Equal(t, curBytes, int(totalBytes))
		assert.Equal(t, int(buffer[4]), SET)
		assert.Equal(t, int(buffer[5]), 1)

		break
	}
}

func TestClientGet(t *testing.T) {
	go func() {
		server := Server{"127.0.0.1", "1234"}
		server.Run(4 * 1024 * 1024)
	}()

	<-time.After(2 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	assert.Equal(t, err, nil)

	setBuffer := makeSetBytes("name", []byte("jerry"))
	wn, wok := conn.Write(setBuffer)
	assert.Equal(t, wok, nil)
	assert.Equal(t, wn, len(setBuffer))

	buffer := []byte{}
	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, err, nil)

		t.Logf("Recv %d bytes from server", n)

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		t.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

		if curBytes < int(totalBytes) {
			continue
		}

		assert.Equal(t, curBytes, int(totalBytes))
		assert.Equal(t, int(buffer[4]), SET)
		assert.Equal(t, int(buffer[5]), 1)

		buffer = buffer[totalBytes:]

		break
	}

	// request server for get `name`.
	getBuffer := makeGetBytes("name")
	wn, wok = conn.Write(getBuffer)
	assert.Equal(t, wok, nil)
	assert.Equal(t, wn, len(getBuffer))

	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, err, nil)

		t.Logf("Recv %d bytes from server", n)

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		t.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

		if curBytes < int(totalBytes) {
			continue
		}

		assert.Equal(t, curBytes, int(totalBytes))
		assert.Equal(t, int(buffer[4]), GET)
		assert.Equal(t, int(buffer[5]), 1)

		bufferCpy := buffer[6:]
		valueBytes := binary.LittleEndian.Uint32(bufferCpy)
		assert.Equal(t, int(valueBytes), 9)
		assert.Equal(t, string(bufferCpy[4:]), "jerry")

		buffer = buffer[totalBytes:]
		break
	}

	// request server for get `NAME`.
	getBuffer = makeGetBytes("NAME")
	wn, wok = conn.Write(getBuffer)
	assert.Equal(t, wok, nil)
	assert.Equal(t, wn, len(getBuffer))

	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, err, nil)

		t.Logf("Recv %d bytes from server", n)

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		t.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

		if curBytes < int(totalBytes) {
			continue
		}

		assert.Equal(t, curBytes, int(totalBytes))
		assert.Equal(t, int(buffer[4]), GET)
		assert.Equal(t, int(buffer[5]), 0)

		buffer = buffer[6:]
		valueBytes := binary.LittleEndian.Uint32(buffer)
		assert.Equal(t, int(valueBytes), 4)

		break
	}
}

func TestClientDelete(t *testing.T) {
	go func() {
		server := Server{"127.0.0.1", "1234"}
		server.Run(4 * 1024 * 1024)
	}()

	<-time.After(2 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	assert.Equal(t, err, nil)

	// request server for set `name:jerry`.
	setBuffer := makeSetBytes("name", []byte("jerry"))
	wn, wok := conn.Write(setBuffer)
	assert.Equal(t, wok, nil)
	assert.Equal(t, wn, len(setBuffer))

	buffer := []byte{}
	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, err, nil)

		t.Logf("Recv %d bytes from server", n)

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		t.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

		if curBytes < int(totalBytes) {
			continue
		}

		assert.Equal(t, curBytes, int(totalBytes))
		assert.Equal(t, int(buffer[4]), SET)
		assert.Equal(t, int(buffer[5]), 1)

		buffer = buffer[totalBytes:]
		break
	}

	// request server for get `name`.
	getBuffer := makeGetBytes("name")
	wn, wok = conn.Write(getBuffer)
	assert.Equal(t, wok, nil)
	assert.Equal(t, wn, len(getBuffer))

	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, err, nil)

		t.Logf("Recv %d bytes from server", n)

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		t.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

		if curBytes < int(totalBytes) {
			continue
		}

		assert.Equal(t, curBytes, int(totalBytes))
		assert.Equal(t, int(buffer[4]), GET)
		assert.Equal(t, int(buffer[5]), 1)

		bufferCpy := buffer[6:]
		valueBytes := binary.LittleEndian.Uint32(bufferCpy)
		assert.Equal(t, int(valueBytes), 9)
		assert.Equal(t, string(bufferCpy[4:]), "jerry")

		buffer = buffer[totalBytes:]
		break
	}

	// request server for delete `name`.
	setBuffer = makeDeleteBytes("name")
	wn, wok = conn.Write(setBuffer)
	assert.Equal(t, wok, nil)
	assert.Equal(t, wn, len(setBuffer))

	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, err, nil)

		t.Logf("Recv %d bytes from server", n)

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		t.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

		if curBytes < int(totalBytes) {
			continue
		}

		assert.Equal(t, curBytes, int(totalBytes))
		assert.Equal(t, int(buffer[4]), DEL)
		assert.Equal(t, int(buffer[5]), 1)

		buffer = buffer[totalBytes:]
		break
	}

	// request server for get `name` again, `name:jerry` is deleted in previous.
	getBuffer = makeGetBytes("name")
	wn, wok = conn.Write(getBuffer)
	assert.Equal(t, wok, nil)
	assert.Equal(t, wn, len(getBuffer))

	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, err, nil)

		t.Logf("Recv %d bytes from server", n)

		buffer = append(buffer, readBuf[:n]...)

		if n == 0 || len(buffer) < 4 {
			continue
		}

		totalBytes := binary.LittleEndian.Uint32(buffer)
		curBytes := len(buffer)
		t.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

		if curBytes < int(totalBytes) {
			continue
		}

		assert.Equal(t, curBytes, int(totalBytes))
		assert.Equal(t, int(buffer[4]), GET)
		assert.Equal(t, int(buffer[5]), 0)

		buffer = buffer[6:]
		valueBytes := binary.LittleEndian.Uint32(buffer)
		assert.Equal(t, int(valueBytes), 4)

		break
	}
}

// following are benchmark functions.

func BenchmarkSet(b *testing.B) {
	b.StopTimer()
	go func() {
		server := Server{"127.0.0.1", "1234"}
		server.Run(4 * 1024 * 1024)
	}()

	<-time.After(1 * time.Second)
	b.StartTimer()
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i++
			conn, err := net.Dial("tcp", "127.0.0.1:1234")
			assert.Equal(b, err, nil)

			keyStr := "name" + strconv.Itoa(i)
			setBuffer := makeSetBytes(keyStr, []byte("jerry"))
			b.Logf(keyStr)
			wn, wok := conn.Write(setBuffer)
			assert.Equal(b, wok, nil)
			assert.Equal(b, wn, len(setBuffer))

			buffer := []byte{}
			for {
				readBuf := make([]byte, 256)
				n, err := bufio.NewReader(conn).Read(readBuf)
				assert.Equal(b, err, nil)

				b.Logf("Recv %d bytes from server", n)

				buffer = append(buffer, readBuf[:n]...)

				if n == 0 || len(buffer) < 4 {
					continue
				}

				totalBytes := binary.LittleEndian.Uint32(buffer)
				curBytes := len(buffer)
				b.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

				if curBytes < int(totalBytes) {
					continue
				}

				assert.Equal(b, curBytes, int(totalBytes))
				assert.Equal(b, int(buffer[4]), SET)
				assert.Equal(b, int(buffer[5]), 1)

				buffer = buffer[totalBytes:]

				break
			}
		}
	})

	// for i := 0; i < b.N; i++ {
	// 	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	// 	assert.Equal(b, err, nil)

	// 	keyStr := "name" + strconv.Itoa(i)
	// 	setBuffer := makeSetBytes(keyStr, []byte("jerry"))
	// 	b.Logf(keyStr)
	// 	wn, wok := conn.Write(setBuffer)
	// 	assert.Equal(b, wok, nil)
	// 	assert.Equal(b, wn, len(setBuffer))

	// 	buffer := []byte{}
	// 	for {
	// 		readBuf := make([]byte, 256)
	// 		n, err := bufio.NewReader(conn).Read(readBuf)
	// 		assert.Equal(b, err, nil)

	// 		b.Logf("Recv %d bytes from server", n)

	// 		buffer = append(buffer, readBuf[:n]...)

	// 		if n == 0 || len(buffer) < 4 {
	// 			continue
	// 		}

	// 		totalBytes := binary.LittleEndian.Uint32(buffer)
	// 		curBytes := len(buffer)
	// 		b.Logf("Recv progress %d/%d\n", curBytes, totalBytes)

	// 		if curBytes < int(totalBytes) {
	// 			continue
	// 		}

	// 		assert.Equal(b, curBytes, int(totalBytes))
	// 		assert.Equal(b, int(buffer[4]), SET)
	// 		assert.Equal(b, int(buffer[5]), 1)

	// 		buffer = buffer[totalBytes:]

	// 		break
	// 	}
	// }
}
