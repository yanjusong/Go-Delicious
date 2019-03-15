package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt1 := intSeq()

	fmt.Println(nextInt1())
	fmt.Println(nextInt1())
	fmt.Println(nextInt1())

	newInt2 := intSeq()
	fmt.Println(newInt2())
	fmt.Println(newInt2())

	newInt2 = intSeq()
	fmt.Println(newInt2())
}
