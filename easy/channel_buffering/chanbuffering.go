package main

import "fmt"

func main() {
	msg := make(chan string, 2)

	msg <- "Hello"
	msg <- "World"
	// 超过缓冲大小会阻塞死锁
	// msg <- "goland"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
	// 缓冲区为空，阻塞死锁
	// fmt.Println(<-msg)
}
