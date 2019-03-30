package main

import "fmt"

type student struct {
	name string
	id   int
}

func (stu *student) increaseId() {
	stu.id++
}

func (stu student) decreaseId() {
	stu.id--
}

// NOTE:
// What difference between increadeId and decreaseId?
// You might easily see that the declaration style is different,
// the increadeId is using a pointer style refer to `stu *student`,
// another is not.
// The function declaration using pointer will change the actual variable,
// call increaseId, the id in student's field will be changed.
// The function declaration using value type is just copy a instance of student,
// of course the id field is changed in decreaseId scope,
// but the caller instance filed will not be changed.
// Above comment is working, whether the caller always as an variable of student
// is an pointer or is an value.

func main() {
	stu1 := student{
		name: "tom",
		id:   1,
	}
	fmt.Println("origin stu1:", stu1)
	stu1.increaseId()
	fmt.Println("after indreaseId, stu1:", stu1)
	stu1.decreaseId()
	fmt.Println("after decreaseId, stu1:", stu1)
	fmt.Println(stu1)

	fmt.Println()

	stu2 := &student{
		name: "tom",
		id:   1,
	}
	fmt.Println("origin stu2:", stu2)
	stu2.increaseId()
	fmt.Println("after indreaseId, stu2:", stu2)
	stu2.decreaseId()
	fmt.Println("after decreaseId, stu2:", stu2)
	fmt.Println(stu2)
}
