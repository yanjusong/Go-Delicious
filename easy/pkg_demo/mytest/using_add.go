package mytest

import (
	"Gogoyan/easy/pkg_demo/mymath"
	"fmt"
)

func init() {
	fmt.Println("package: mytest, file: using_add.go, position: init1")
}

func init() {
	fmt.Println("package: mytest, file: using_add.go, position: init2")
}

func PrintAdd(a, b int) {
	sum := mymath.Add(a, b)
	fmt.Printf("%d+%d=%d\n", a, b, sum)
}
