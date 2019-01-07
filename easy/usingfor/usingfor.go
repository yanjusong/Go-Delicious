package main

import (
	"fmt"
)

func print9x9() {
	var i int
	var j int

	for i = 0; i < 9; i++ {
		for j = 0; j < i+1; j++ {
			fmt.Printf("%dx%d=%2d ", j+1, i+1, (j+1)*(i+1))
		}
		fmt.Printf("\n")
	}
}

func binarySearch(nums []int, value int) int {
	left := 0
	right := len(nums) - 1
	var mid int

	for left <= right {
		mid = (left + right) / 2
		if value < nums[mid] {
			right = mid - 1
		} else if value > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func test_binary_search() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0

	fmt.Println("nums: ", nums)

	for i = -1; i < 12; i++ {
		ret := binarySearch(nums, i)
		fmt.Printf("search[%d]: result:%d\n", i, ret)
	}
}

func main() {
	print9x9()
	test_binary_search()
}
