package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// go显式转换方法：
// 表达式 T(v) 将值 v 转换为类型 T

func main() {
	fmt.Println("case 1:")
	{
		var bytevar byte = 100
		fmt.Println(bytevar)

		// var intvar int = bytevar // error:不支持隐式转换
		var intvar int = int(bytevar) // ok:显式转换
		fmt.Println(intvar)

		// var boolvar bool = bool(intvar) // error:不能将int转换到bool

		var a int = 1
		var b int = 3
		var c float32 = float32(a) / float32(b)
		fmt.Println(c)
	}

	fmt.Println("\ncase 2:")
	{
		var i int = 1
		fmt.Printf("i value is %v, type is %v\n", i, reflect.TypeOf(i))

		var f float64 = float64(i)
		fmt.Printf("f value is %v, type is %v\n", f, reflect.TypeOf(f))

		var ui uint = uint(i)
		fmt.Printf("ui value is %v, type is %v\n", ui, reflect.TypeOf(ui))
	}

	fmt.Println("\ncase 3:")
	{
		// atoi
		i, ok := strconv.ParseInt("1000", 10, 0)
		if ok == nil {
			fmt.Printf("ParseInt , i is %v , type is %v\n", i, reflect.TypeOf(i))
		}

		ui, ok := strconv.ParseUint("100", 10, 0)
		if ok == nil {
			fmt.Printf("ParseUint , ui is %v , type is %v\n", ui, reflect.TypeOf(i))
		}

		oi, ok := strconv.Atoi("100")
		if ok == nil {
			fmt.Printf("Atoi , oi is %v , type is %v\n", oi, reflect.TypeOf(i))
		}

		// itoa
		var j int
		j = 0x100
		str := strconv.FormatInt(int64(j), 10) // FormatInt第二个参数表示进制，10表示十进制。
		fmt.Println(str)
		fmt.Println(reflect.TypeOf(str))
		str = strconv.FormatInt(int64(j), 16) // FormatInt第二个参数表示进制，10表示十进制。
		fmt.Println(str)
		fmt.Println(reflect.TypeOf(str))
	}
}
