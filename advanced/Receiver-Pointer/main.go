package main

import "fmt"

type test struct {
	name string
}

func (t test) TestValue() {
	fmt.Printf("in TestValue: %p\n", &t)
}

func (t *test) TestPointer() {
	fmt.Printf("in TestPointer: %p\n", t)
}

func test1() {
	fmt.Println("in test1 --------------++++++++--------------")
	t := test{}
	fmt.Printf("in main: %p\n", &t) // 0xc00000e200

	t.TestValue() // 0xc00000e210
	// t非也能用类型，`t.TestValue()`相当于下面的转换兵调用：
	vFunc := test.TestValue
	vFunc(t)

	t.TestPointer() // 0xc00000e200
	// t非也能用类型，`t.TestPointer()`相当于下面的转换兵调用：
	pFunc := (*test).TestPointer
	pFunc(&t)
}

func test2() {
	fmt.Println("in test2 --------------++++++++--------------")
	t := &test{}
	fmt.Printf("in main: %p\n", t) // 0xc00000c030

	t.TestValue() // 0xc00000e230
	// t为引用类型，`t.TestValue()`相当于下面的转换兵调用：
	vFunc := test.TestValue
	vFunc(*t)

	t.TestPointer() // 0xc00000e220
	// t为引用类型，`t.TestPointer()`相当于下面的转换兵调用：
	pFunc := (*test).TestPointer
	pFunc(t)
}

func main() {
	test1()
	test2()
}
