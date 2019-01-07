package main

import "fmt"

func increase(p *int) {
	*p++
}

func printVarAddr() {
	var i int
	var pi *int

	fmt.Printf("address of i is: 0x%x\n", &i)
	pi = &i
	fmt.Printf("value of pi is: 0x%x\n", pi)

	i = 100

	fmt.Printf("i=%d, *pi=%d\n", i, *pi)

	fmt.Printf("after call increase(&x): ")
	increase(&i)
	fmt.Printf("i=%d, *pi=%d\n", i, *pi)

	fmt.Printf("after call increase(pi): ")
	increase(pi)
	fmt.Printf("i=%d, *pi=%d\n", i, *pi)
}

func testPointerArray() {
	nums := [3]int{1, 2, 3}
	var p [len(nums)]*int

	for i := 0; i < len(nums); i++ {
		p[i] = &nums[i]
	}

	for i := 0; i < len(p); i++ {
		fmt.Printf("nums[%d]:%d\n", i, *p[i])
	}

}

func testPointer2Pointer() {
	var i int
	var pi *int
	var ppi **int

	i = 99
	pi = &i
	ppi = &pi

	fmt.Printf("i=%d\n", i)
	fmt.Printf("*pi=%d\n", *pi)
	fmt.Printf("*ppi=%d\n", **ppi)
}

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func testSwap() {
	x := 1
	y := 2
	fmt.Printf("x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("after swap: x=%d, y=%d\n", x, y)
}

func main() {
	printVarAddr()
	testPointerArray()
	testPointer2Pointer()
	testSwap()
}
