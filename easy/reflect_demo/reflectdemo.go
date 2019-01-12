package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}

func (u User) Foo() {
	// Do nothing
}

func main() {
	u := User{"张三", 20}

	// 获取对象的具体类型
	t := reflect.TypeOf(u)
	fmt.Println(t)

	// 反射获取一个对象的reflect.Value
	v := reflect.ValueOf(u)
	fmt.Println(v)
	vt := reflect.TypeOf(v)
	fmt.Println(vt)

	fmt.Printf("---------------------------------------------\n")
	// 打印类型
	fmt.Printf("%T\n", u)
	// 打印值
	fmt.Printf("%v\n", u)

	fmt.Printf("---------------------------------------------\n")
	// reflect.Value转成具体类型
	u1 := v.Interface().(User)
	fmt.Println(u1)
	ut1 := reflect.TypeOf(u1)
	fmt.Println(ut1)
	t1 := v.Type()
	fmt.Println(t1)

	fmt.Println("\n底层类型: ", t.Kind())

	fmt.Printf("\nmember in struct User:\n")
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	fmt.Printf("\nfunction in struct User:\n")
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}

	fmt.Printf("---------------------------------------------\n")
	mPrint := v.MethodByName("Print")
	// 转化为[]byte
	args := []reflect.Value{reflect.ValueOf("前缀")}

	fmt.Println(mPrint.Call(args))
}
