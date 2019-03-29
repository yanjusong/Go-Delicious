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

func setTriggle(t testing.TB, conn net.Conn, key string, value []byte, buffer []byte) {
	setBuffer := makeSetBytes(key, value)
	wn, wok := conn.Write(setBuffer)
	assert.Equal(t, nil, wok)
	assert.Equal(t, len(setBuffer), wn)

	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, nil, err)

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

		assert.Equal(t, int(totalBytes), curBytes)
		assert.Equal(t, SET, int(buffer[4]))
		assert.Equal(t, 1, int(buffer[5]))

		buffer = buffer[totalBytes:]
		assert.Equal(t, 0, len(buffer))

		break
	}
}

func delTriggle(t testing.TB, conn net.Conn, key string, buffer []byte) {
	setBuffer := makeDeleteBytes(key)
	wn, wok := conn.Write(setBuffer)
	assert.Equal(t, nil, wok)
	assert.Equal(t, len(setBuffer), wn)

	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, nil, err)

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

		assert.Equal(t, int(totalBytes), curBytes)
		assert.Equal(t, DEL, int(buffer[4]))
		assert.Equal(t, 1, int(buffer[5]))

		buffer = buffer[totalBytes:]
		assert.Equal(t, 0, len(buffer))

		break
	}
}

func getTriggle(t testing.TB, conn net.Conn, key string, expectedStat int, expectedStrLen int, expectedStr string, buffer []byte) {
	getBuffer := makeGetBytes(key)
	wn, wok := conn.Write(getBuffer)
	assert.Equal(t, nil, wok)
	assert.Equal(t, len(getBuffer), wn)

	for {
		readBuf := make([]byte, 256)
		n, err := bufio.NewReader(conn).Read(readBuf)
		assert.Equal(t, nil, err)

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

		assert.Equal(t, int(totalBytes), curBytes)
		assert.Equal(t, GET, int(buffer[4]))
		assert.Equal(t, expectedStat, int(buffer[5]))

		bufferCpy := buffer[6:]
		valueBytes := binary.LittleEndian.Uint32(bufferCpy)
		assert.Equal(t, 4+expectedStrLen, int(valueBytes))

		if expectedStat == 1 {
			assert.Equal(t, expectedStr, string(bufferCpy[4:]))
		}

		buffer = buffer[totalBytes:]
		assert.Equal(t, 0, len(buffer))
		break
	}
}

func TestClientSet(t *testing.T) {
	go func() {
		server := Server{"127.0.0.1", "1234"}
		server.Run(4 * 1024 * 1024)
	}()

	<-time.After(2 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	assert.Equal(t, nil, err)

	buffer := []byte{}
	setTriggle(t, conn, "name", []byte("jerry"), buffer)
	setTriggle(t, conn, "", []byte("jerry"), buffer)
	setTriggle(t, conn, "", []byte(""), buffer)
	setTriggle(t, conn, "addr", []byte(""), buffer)
}

func TestClientGet(t *testing.T) {
	go func() {
		server := Server{"127.0.0.1", "1234"}
		server.Run(4 * 1024 * 1024)
	}()

	<-time.After(2 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	assert.Equal(t, nil, err)

	buffer := []byte{}
	setTriggle(t, conn, "name", []byte("jerry"), buffer)
	getTriggle(t, conn, "name", 1, 5, "jerry", buffer)
	getTriggle(t, conn, "NAME", 0, 0, "", buffer)
}

func TestClientDelete(t *testing.T) {
	go func() {
		server := Server{"127.0.0.1", "1234"}
		server.Run(4 * 1024 * 1024)
	}()

	<-time.After(2 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	assert.Equal(t, nil, err)

	buffer := []byte{}
	setTriggle(t, conn, "name", []byte("jerry"), buffer)
	getTriggle(t, conn, "name", 1, 5, "jerry", buffer)
	delTriggle(t, conn, "name", buffer)
	getTriggle(t, conn, "name", 0, 0, "", buffer)
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
		buffer := []byte{}

		for pb.Next() {
			i++
			conn, err := net.Dial("tcp", "127.0.0.1:1234")
			assert.Equal(b, nil, err)

			keyStr := "name" + strconv.Itoa(i)

			setTriggle(b, conn, keyStr, []byte("jerry"), buffer)
			getTriggle(b, conn, keyStr, 1, 5, "jerry", buffer)
		}
	})
}
