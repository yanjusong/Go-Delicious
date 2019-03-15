package main

import "fmt"

func main() {
	fmt.Println("case 1:")
	{
		nums := [3]int{1, 2, 3}

		// 针对数组的range是一个复制，
		// 当i==0是对第二个元素修改，
		// 由于range的复制在对元素修改之前，
		// 所以访问到第二个元素还是之前的数据。
		for i, x := range nums {
			if i == 0 {
				nums[1] = 200
			}

			fmt.Println(i, x)
		}

		fmt.Println(nums)
	}

	fmt.Println("\ncase 2:")
	{
		nums := []int{1, 2, 3}

		// 引用类型的range类似于指针的传递
		// 和数组的复制有明显区别
		for i, x := range nums {
			if i == 0 {
				nums[1] = 200
			}

			fmt.Println(i, x)
		}

		fmt.Println(nums)
	}
}
