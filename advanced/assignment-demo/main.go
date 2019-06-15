package main

import "fmt"

type Easy struct {
	name   string
	age    int
	scores map[string]int
}

func main() {
	e1 := Easy{
		name: "Jason",
		age:  24,
		scores: map[string]int{
			"Math": 95,
		},
	}
	// 浅拷贝，类似于C++中的拷贝，go里面引用类型会有共同副本。
	e2 := e1
	fmt.Println("e1:", e1)
	fmt.Println("e2:", e2)

	e1.age++
	fmt.Println("after e1.age++")
	fmt.Println("e1:", e1)
	fmt.Println("e2:", e2)

	e2.scores["English"] = 88
	fmt.Println("after add English score in e2")
	fmt.Println("e1:", e1)
	fmt.Println("e2:", e2)
}
