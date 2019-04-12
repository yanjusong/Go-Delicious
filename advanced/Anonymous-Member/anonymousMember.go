package main

import "fmt"

type anier interface {
	foo()
	bar()
}

type Base struct {
	iBase   int
	strBase string
}

func (base Base) foo() {
	fmt.Println("In Base::foo")
}

func (base *Base) bar() {
	fmt.Println("In Base::bar")
}

type A struct {
	Base
	iA   int
	strA string
}

type B struct {
	*Base
	iB   int
	strB string
}

func callFoo(a anier) {
	a.foo()
	a.bar()
}

// NOTE:
// 1.类型 T 方法集包含全部 receiver T 方法。
// 2.类型 *T 方法集包含全部 receiver T + *T 方法。
// 3.如类型 S 包含匿名字段 T，则 S 方法集包含 T 方法, *S 方法集包含 T + *T 方法.
// 4.如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
// 5.不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。（总结3、4）
func main() {
	var i anier

	base := Base{
		iBase:   0,
		strBase: "Base",
	}
	callFoo(&base)
	// callFoo(base) // error，因为规则1

	i = &base
	callFoo(i)
	i.foo()
	i.bar()

	a := A{
		Base: Base{
			iBase:   0,
			strBase: "Base",
		},
		iA:   1,
		strA: "A",
	}
	fmt.Println("\npos 1:")
	fmt.Println(a)
	// callFoo(a) //error，因为规则3
	callFoo(&a)

	pa := &a
	fmt.Println("\npos 2:")
	fmt.Println(pa)
	callFoo(pa)

	b := B{
		Base: &Base{
			iBase:   0,
			strBase: "Base",
		},
		iB:   1,
		strB: "B",
	}
	fmt.Println("\npos 3:")
	fmt.Println(b)
	callFoo(b)

	pb := &b
	fmt.Println("\npos 4:")
	fmt.Println(pb)
	callFoo(pb)
}
