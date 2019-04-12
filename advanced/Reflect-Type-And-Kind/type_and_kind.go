package main

import (
	"fmt"
	"reflect"
)

// MyInt ...
type MyInt int

func main() {
	var x MyInt = 7

	fmt.Println("type:", reflect.TypeOf(x))            // main.MyInt
	fmt.Println("value:", reflect.ValueOf(x))          // value: 7
	fmt.Println("value:", reflect.ValueOf(x).String()) // <main.MyInt Value>

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type()) // <main.MyInt Value>

	// Kind()方法获取底层的类型
	fmt.Println("kind:", v.Kind()) // kind: int

	fmt.Println(v.Interface()) // 7
}
