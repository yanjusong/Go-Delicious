package main

import (
	"fmt"
	"reflect"
)

func getSumStr(s string, n ...int) string {
	// n是一个slice对象
	fmt.Println(reflect.TypeOf(n))

	var sum int
	for _, i := range n {
		sum += i
	}

	return fmt.Sprintf(s, sum)
}

func multipleReturn() (int, int, int) {
	return 1, 2, 3
}

func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}

	return x
}

type WorkFunc func(a, b int) string

func strategyFunc(f WorkFunc, a, b int) string {
	return f(a, b)
}

func add(a, b int) string {
	res := a + b
	resStr := fmt.Sprintf("%d+%d=%d\n", a, b, res)
	return resStr
}

func sub(a, b int) (resStr string) {
	res := a - b
	resStr = fmt.Sprintf("%d-%d=%d\n", a, b, res)
	return
}

func increaseDefer(i int) (x int) {
	defer func() {
		x++
		fmt.Println("in defer 1, x=", x)
	}()

	x = i + 1

	// defer函数的调用顺序和注册顺序相反
	defer func() {
		x++
		fmt.Println("in defer 2, x=", x)
	}()

	if true {
		return
	}

	// return之后的defer不会被调用
	defer func() {
		x++
		fmt.Println("in defer 3, x=", x)
	}()

	return
}

func main() {
	fmt.Println("case 1:")
	{
		s := []int{1, 2, 3}

		// 将slice当作多参数应该如下调用
		res := getSumStr("sum: %d", s...)
		println(res)

		res = getSumStr("sum: %d", 1, 2, 3, 4)
		println(res)
	}

	fmt.Println("\ncase 2:")
	{
		// strategyFunc函数的一个参数是函数类型
		a, b := 1, 2
		res1 := strategyFunc(add, a, b)
		fmt.Println(res1)

		res2 := strategyFunc(sub, a, b)
		fmt.Println(res2)

		// 匿名函数当作参数
		res3 := strategyFunc(func(a, b int) string {
			res := a * b
			resStr := fmt.Sprintf("%d*%d=%d\n", a, b, res)
			return resStr
		}, a, b)
		fmt.Println(res3)
	}

	fmt.Println("\ncase 3:")
	{
		// error
		// res := getSumStr("sum: %d", multipleReturn())
		fmt.Println(sum(multipleReturn()))
	}

	fmt.Println("\ncase 4:")
	{
		x := increaseDefer(1)
		fmt.Println(x)
	}
}
