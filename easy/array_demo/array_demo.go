package main

import (
	"fmt"
)

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

func change(nums [3]int) {
	fmt.Printf("%p\n", &nums)
	nums[1] = 100
	fmt.Println("in change() nums:", nums)
}

func main() {
	fmt.Println("case 1:")
	{
		a := [3]int{1, 2}           // 未初始化元素值为 0。
		b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
		c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
		d := [...]struct {
			name string
			age  uint8
		}{
			{"user1", 10}, // 可省略元素类型。
			{"user2", 20}, // 别忘了最后一行的逗号。
		}
		fmt.Println(arr0, arr1, arr2, str)
		fmt.Println(a, b, c, d)
	}

	fmt.Println("\ncase 2:")
	{
		a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
		b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
		fmt.Println(arr0, arr1)
		fmt.Println(a, b)

		// 遍历二维数组
		for i, nums := range b {
			for j, x := range nums {
				fmt.Printf("(%d,%d)=%d ", i, j, x)
			}
			fmt.Println()
		}
	}

	fmt.Println("\ncase 3:")
	{
		nums := [3]int{1, 2, 3}
		fmt.Println("in main() nums:", nums)
		fmt.Printf("%p\n", &nums)

		// 数组为参数是传值，复制一份过去
		change(nums)
		fmt.Println("after call change(), in main() nums:", nums)
	}
}
