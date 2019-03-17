package main

import "fmt"

type Sayer interface {
	Say(message string)
	SayHi()
}

type Animal struct {
	name string
}

// 继承Animal
type Dog struct {
	Animal
	id int
}

func (a *Animal) Say(message string) {
	fmt.Printf("Animal[%v] say: %v\n", a.name, message)
}

func (a *Animal) SayHi() {
	// 会调用Animal的Say，尽管传入的是指向Dog类的对象
	a.Say("Hi")
}

// 重写Say(message string)
func (d *Dog) Say(message string) {
	fmt.Printf("Dog[%v] say: %v\n", d.name, message)
}

// for case4
/////////////////////////////////////////////
type Stringer interface {
	String() string
}

// Printer是Stringer的超集
type Printer interface {
	String() string
	Print()
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %v", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func main() {
	fmt.Println("case 1:")
	{
		var sayer Sayer = &Dog{Animal{name: "Jason"}, 1}

		sayer.Say("hello world") // Dog[Yoda] say: hello world
		sayer.SayHi()            // Animal[Yoda] say: Hi

		var a *Animal = &Animal{"Jerry"}
		a.SayHi()
		a.Say("--->")

		var d *Dog = &Dog{Animal{name: "Tom"}, 2}
		d.SayHi()
		d.Say("--->")
	}

	fmt.Println("\ncase 2:")
	{
		v := Dog{Animal{name: "Jason"}, 1}
		p := &Dog{Animal{name: "Jeff"}, 2}
		fmt.Printf("Type: %T, Value: %v\n", v, v)
		fmt.Printf("Type: %T, Value: %v\n", p, p)

		fmt.Println("after id increase:")
		v.id = v.id + 1
		p.id = p.id + 1
		fmt.Printf("Type: %T, Value: %v\n", v, v)
		fmt.Printf("Type: %T, Value: %v\n", p, p)

		var vi interface{} = v
		var pi interface{} = &v
		// vi.(Dog).id = 10 //error: can't assign
		pi.(*Dog).id = 10

		fmt.Printf("%v\n", vi.(Dog))
		fmt.Printf("%v\n", pi.(*Dog))
		fmt.Printf("%v\n", vi)
		fmt.Printf("%v\n", pi)
	}

	fmt.Println("\ncase 3:")
	{
		var o interface{} = &Dog{Animal{name: "Jeff"}, 2}

		if i, ok := o.(fmt.Stringer); ok {
			fmt.Println(i)
		} else {
			fmt.Println("convert to fmt.Stringer error")
		}

		if i, ok := o.(*Dog); ok {
			fmt.Println(i)
		} else {
			fmt.Println("convert to *Dog error")
		}

		d := o.(*Dog)
		d.id = 100
		fmt.Printf("Type: %T, Value: %v\n", d, d)

		switch v := o.(type) {
		case nil: // o == nil
			fmt.Println("nil")
		case fmt.Stringer: // interface
			fmt.Println(v)
		case func() string: // func
			fmt.Println(v())
		case *Dog: // *struct
			fmt.Printf("%d, %s\n", v.id, v.name)
			// 这样写也可以
			fmt.Printf("%d, %s\n", v.id, v.Animal.name)
		default:
			fmt.Println("unknown")
		}
	}

	fmt.Println("\ncase 4:")
	{
		var o Printer = &User{1, "Tom"}
		var s Stringer = o
		fmt.Println(s.String())

		// error: 超集接口对象可转换为子集接口，反之出错。
		// var o Stringer = &User{1, "Tom"}
		// var s Printer = o
		// fmt.Println(s.String())
	}
}
