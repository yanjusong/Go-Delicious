package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	SET      = 1
	GET byte = 2
	DEL byte = 3
	ERR byte = 50
)

func doNothing(buffer []byte) {
	buffer[0] = 200
	buffer = buffer[3:4]
}

func main() {
	nums := []byte{1, 2, 3, 4}
	doNothing(nums)
	fmt.Println(nums)
	nums2 := []byte{}
	nums2 = append(nums2, nums[2:3]...)
	fmt.Println(nums2)
	nums2[0] = 100
	fmt.Println(nums)
	fmt.Println(nums2)

	str := []byte("hello")
	strs := string(str)
	fmt.Println(str)
	fmt.Println(strs)

	str[2] = 'L'
	fmt.Println(str)
	fmt.Println(strs)
	return
	// fmt.Println(nslice[0])
	// fmt.Println(nslice[1])
	// fmt.Println(nslice[2])
	// fmt.Println(nslice[3])

	// data := binary.LittleEndian.Uint32(nslice)
	// fmt.Println(data)

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	var n uint32 = 100
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)

		nslice := make([]byte, 4)
		nslice[0] = byte(n & 0x000000ff)
		nslice[1] = byte((n & 0x0000ff00) >> 8)
		nslice[2] = byte((n & 0x00ff0000) >> 16)
		nslice[3] = byte((n & 0xff000000) >> 24)

		wn, _ := conn.Write(nslice)
		fmt.Println("write to server, bytes:", wn)
		n++

		// send to socket
		// fmt.Fprintf(conn, text+"\n")
		// listen for reply
		// message, _ := bufio.NewReader(conn).ReadString('\n')
		// fmt.Print("Message from server: " + message)
	}
}
