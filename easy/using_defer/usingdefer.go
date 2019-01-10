package main

import "fmt"

func foo1() {
	fmt.Printf("In foo1()\n")
}

func foo() {
	fmt.Printf("In foo()\n")
	defer func() {
		fmt.Printf("inter defer in foo()\n")
	}()
}

func bar1() (result int) {
	defer func() {
		result++
	}()
	result = 1

	return result
}

func bar2(i *int) {
	defer func() {
		*i++
	}()

	*i = 100
}

func main() {
	fmt.Printf("Hello\n")

	defer foo1()

	foo()
	fmt.Printf("World\n")

	defer func() {
		fmt.Printf("inter defer in main()\n")
	}()

	bar1Res := bar1()
	fmt.Printf("bar1Res=%d\n", bar1Res)

	var x int
	bar2(&x)
	fmt.Printf("x=%d\n", x)
}
