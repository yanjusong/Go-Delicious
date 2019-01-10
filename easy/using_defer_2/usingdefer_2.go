package main

import "fmt"

// 1. defer调用的函数参数会立马赋值
func testDefer1() {
	i := 0
	defer fmt.Printf("i=%d\n", i)
	i++
}

// 2. defer调用的函数可以改变返回值
func testDefer2() (result int) {
	defer func() {
		result++
	}()

	return 0
}

// 3. defer会将函数放入栈中，调用时后进先出
func testDefer3() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("in testDefer3() i=%d\n", i)
		}()
	}

	for i := 0; i < 5; i++ {
		defer fmt.Printf("in testDefer3() i=%d\n", i)
	}
}

func main() {
	testDefer1()
	ret := testDefer2()
	fmt.Printf("ret=%d\n", ret)
	testDefer3()
}
