package main

import (
	"fmt"
	"unsafe"
)

// `:=`只能在函数内使用；全局变量可以定义且不使用，局部变量不可以。
// g := 1 // error
var g int = 1 // ok

// 这个传参的过程类似于C语言
func change(nums []int) {
	fmt.Printf("slice addr: %p\n", nums)
	fmt.Printf("var addr: %p\n", &nums)
	nums = append(nums, 100)
	fmt.Printf("slice addr: %p\n", nums)
	fmt.Printf("var addr: %p\n", &nums)
}

func main() {
	fmt.Println("case 1:")
	{
		nums := [3]int{1, 2, 3}
		i := 1
		fmt.Println(i, nums)

		// 多变量赋值时，先计算所有相关值，再依次从左到右赋值
		i, nums[i] = 2, 100
		fmt.Println(i, nums)

		i, nums[i] = 200, 100
		fmt.Println(i, nums)
	}

	fmt.Println("\ncase 2:")
	{
		nums := []int{1, 2, 3}
		fmt.Printf("slice addr: %p\n", nums)
		fmt.Printf("var addr: %p\n", &nums)
		change(nums)
	}

	fmt.Println("\ncase 3:")
	{
		// 未使用的局部常量不会引发编译错误

		// const a := 1 // error: 常量不能使用 ":=" 语法定义。
		const a int = 1 // ok

		const (
			b = 10
			c = 100
		)

		const d, e = 9, "eeeee"
		// const f int // error
		const f, g int = 9, 99

		const (
			s1 = "hello"
			s2 // 在常量组中，如不提供类型和初始化值，那么视作与上一个常量相同。
			s3 = "world"
		)

		fmt.Println(s1, s2, s3)

		// 常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值。
		const (
			s         = "hello world"
			lens      = len(s)
			sizeofvar = unsafe.Sizeof(lens)
		)
		fmt.Println(s, lens, sizeofvar)
	}

	fmt.Println("\ncase 4:")
	{
		const (
			Sunday = iota
			Monday
			Tuesday
			Wednesday
			Thursday
			Friday
			Saturday
		)
		fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

		const (
			_        = iota
			KB int64 = 1 << (10 * iota) // iota = 1
			MB                          // iota = 2
			GB                          // iota = 3
			TB                          // iota = 4
		)
		fmt.Println(KB, MB, GB, TB)

		const (
			A, B = iota, iota << 10
			C, D
		)
		fmt.Println(A, B, C, D)
	}
}
