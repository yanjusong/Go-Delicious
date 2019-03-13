package mytest

import (
	"Gogoyan/easy/pkg_demo/mymath"
	"fmt"
)

func init() {
	fmt.Println("package: mytest, file: using_sub.go, position: init1")
}

func init() {
	fmt.Println("package: mytest, file: using_sub.go, position: init2")
}

func PrintSub(a, b int) {
	sum := mymath.Sub(a, b)
	fmt.Printf("%d-%d=%d\n", a, b, sum)
}
