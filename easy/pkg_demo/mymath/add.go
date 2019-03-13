package mymath

import "fmt"

func init() {
	fmt.Println("package: mymath, file: add.go, position: init1")
}

func init() {
	fmt.Println("package: mymath, file: add.go, position: init2")
}

func Add(a, b int) int {
	return a + b
}
