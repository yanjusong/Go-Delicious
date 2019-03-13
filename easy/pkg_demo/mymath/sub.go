package mymath

import "fmt"

func init() {
	fmt.Println("package: mymath, file: sub.go, position: init1")
}

func init() {
	fmt.Println("package: mymath, file: sub.go, position: init2")
}

func Sub(a, b int) int {
	return a - b
}
