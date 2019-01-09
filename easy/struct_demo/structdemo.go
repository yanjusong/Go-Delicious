package main

import "fmt"

type Person struct {
	name string
	age  int
	addr string
}

func printPerson(person *Person) {
	if person == nil {
		fmt.Printf("params is nil.\n")
		return
	}
	fmt.Printf("name:%s\n", person.name)
	fmt.Printf("age:%d\n", person.age)
	fmt.Printf("addr:%s\n", person.addr)
}

func main() {
	var person1 Person
	printPerson(&person1)
	printPerson(nil)

	person2 := Person{"tom", 20, "Beijing"}
	printPerson(&person2)

	person3 := Person{name: "jerry", addr: "Shanghai", age: 30}
	printPerson(&person3)

	fmt.Println(person3)
}
