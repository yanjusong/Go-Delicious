package main

import (
	"fmt"
	"time"
)

func addItems(c chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("add %d\n", i)
		c <- i
	}
}

func removeItems(c <-chan int) {
	for i := range c {
		fmt.Printf("remove %d\n", i)
	}
}

func main() {
	c := make(chan int)

	go addItems(c)
	go removeItems(c)

	time.Sleep(3 * time.Second)
}
