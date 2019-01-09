package main

import "fmt"

func printSliceInfo(s []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(s), cap(s), s)
}

func test_slice() {
	s1 := []int{1, 2, 3}
	printSliceInfo(s1)
	s2 := s1[:]
	printSliceInfo(s2)
	s3 := make([]int, 6)
	printSliceInfo(s3)
	s4 := make([]int, 6, 10)
	printSliceInfo(s4)
	var s5 []int
	printSliceInfo(s5)
	if s5 == nil {
		fmt.Printf("s5 is nil\n")
	}
}

func test_append() {
	s1 := []int{1, 2, 3}
	printSliceInfo(s1)
	s1 = append(s1, 4)
	printSliceInfo(s1)
	s1 = append(s1, 5, 6, 7)
	printSliceInfo(s1)
}

func test_copy() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1), cap(s1)*2)
	printSliceInfo(s2)
	copy(s2, s1)
	printSliceInfo(s2)
}

func main() {
	test_slice()
	test_append()
	test_copy()
}
