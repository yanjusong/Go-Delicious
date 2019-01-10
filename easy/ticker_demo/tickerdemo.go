package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	// ticker开启之后会做重复性的工作

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
