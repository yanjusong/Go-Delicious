package main

import "fmt"

func sum(nums []int, c chan int) {
	result := 0

	for _, x := range nums {
		result += x
	}

	c <- result
}

func testBlockingChan() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9}

	c := make(chan int)

	go sum(nums1, c)
	go sum(nums2, c)

	x, y := <-c, <-c

	fmt.Printf("%d+%d=%d\n", x, y, x+y)
}

func foo(len int, c chan int) {
	for i := 1; i <= len; i++ {
		c <- i
	}

	close(c)
}

func testUnblockingChan() {
	c := make(chan int, 10)
	foo(cap(c), c)

	for x := range c {
		fmt.Printf("%d\n", x)
	}
}

func main() {
	testBlockingChan()
	testUnblockingChan()
}
