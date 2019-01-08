package main

import "fmt"

type Animal interface {
	speak()
	getName() string
}

type Cat struct {
}

type Dog struct {
	name string
}

func (c Cat) speak() {
	fmt.Printf("miao miao miao ~~~\n")
}

func (c Cat) getName() string {
	return "I am cat"
}

func (d Dog) speak() {
	fmt.Printf("wang wang wang ~~~\n")
}

func (d *Dog) getName() string {
	d.name = "旺财"
	return "I am dog"
}

func invoke(p Animal) {
	p.speak()
}

func main() {
	var parent Animal

	c := new(Cat)
	c.speak()

	parent = c
	parent.speak()
	fmt.Printf("parent dese %s\n", parent.getName())

	d := Dog{"Bob"}
	d.speak()
	parent = &d
	parent.speak()
	fmt.Printf("dog name is %s\n", d.name)
	fmt.Printf("parent dese %s\n", parent.getName())
	fmt.Printf("after getName called, dog name is %s\n", d.name)

	fmt.Printf("//////////////////////////////////\n")
	fmt.Println(c)
	fmt.Println(d)
	d2 := new(Dog)
	d2.name = "哈士奇"
	parent = d2
	fmt.Println(d2)

	invoke(c)
	invoke(&d)
	invoke(d2)
}
