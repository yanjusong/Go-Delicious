package main

import "fmt"

func receiver(c1 chan int, c2 chan int, quit chan int) {
	var i, j int
	index := 0

	for {
		index = index + 1
		fmt.Printf("index: %d\n", index)

		// select 会等待一组chan对象中一个或者一个以上可读或者可写
		// 随机从活动的chan对象中选择一个进行操作
		select {
		case i = <-c1:
			fmt.Printf("receive: %d\n", i)
		case j = <-c2:
			fmt.Printf("receive: %d\n", j)
		case <-quit:
			fmt.Printf("quit\n")
			return
		}
	}
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
			c2 <- i + 100
		}
		quit <- 0
	}()

	receiver(c1, c2, quit)
}
