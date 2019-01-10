package main

import "fmt"

func sum(nums ...int) {
	s := 0
	for _, x := range nums {
		s += x
	}

	fmt.Printf("sum=%d\n", s)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
