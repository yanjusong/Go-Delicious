package main

import (
	"fmt"
	"unsafe"
)

// unsafe.Pointer相关：
// 1.任何指针都可以转换为unsafe.Pointer
// 2.unsafe.Pointer可以转换为任何指针
// 3.uintptr可以进行偏移计算

// 函数参数为数组指针，即为指向数组的指针
func change(ptr *[3]int) {
	// 对每个元素都加1
	for i := 0; i < len(*ptr); i++ {
		((*ptr)[i])++
	}
}

func main() {
	fmt.Println("case 1:")
	{
		var i int = 10
		p := &i
		fmt.Println(*p)
		// p++ // error:一般指针不能做加减

		var up uintptr = uintptr(unsafe.Pointer(&i))
		up++ // ok:uinptr可以做偏移运算
	}

	fmt.Println("\ncase 2:")
	{
		u := new(user)
		fmt.Println(*u)

		pName := (*string)(unsafe.Pointer(u))
		*pName = "张三"

		pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
		*pAge = 20

		// 上面相当于：
		// temp := uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)
		// pAge := (*int)(unsafe.Pointer(temp))
		// *pAge = 20

		fmt.Println(*u)
	}

	fmt.Println("\ncase 3:")
	{
		nums := [3]int{1, 2, 3}
		var ptr [3]*int // 指针数组
		for i := 0; i < len(ptr); i++ {
			ptr[i] = &nums[i]
		}

		*ptr[2] = 100

		for i := 0; i < len(ptr); i++ {
			fmt.Println(*ptr[i])
		}
	}

	fmt.Println("\ncase 4:")
	{
		nums := [3]int{1, 2, 3}
		fmt.Println(nums)
		change(&nums)
		fmt.Println(nums)
	}
}

type user struct {
	name string
	age  int
}
