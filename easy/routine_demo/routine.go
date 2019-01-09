package main

import (
	"fmt"
	"time"
)

func foo(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d -> %s\n", i, name)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go foo("tom")
	foo("jerry")
}
