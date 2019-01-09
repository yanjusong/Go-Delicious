package main

import "fmt"

func sum(nums []int) int {
	total := 0

	for i, x := range nums {
		fmt.Printf("add %d\n", nums[i])
		total = total + x
	}

	return total
}

func main() {
	var nums []int

	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}

	s := sum(nums)
	fmt.Printf("Total: %d\n", s)
}
