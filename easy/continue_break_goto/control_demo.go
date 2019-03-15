package main

import "fmt"

func main() {

	//goto直接调到LAbEL2
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LAbEL1
			}
		}
	}
	fmt.Println("PreLAbEL2")
LAbEL1:
	fmt.Println("LastLAbEL2")

	//break跳出和LAbEL1同一级别的循环,继续执行其他的
LAbEL2:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LAbEL2
			}
		}
	}
	fmt.Println("OK")

	//continue
LABEL3:

	for i := 0; i < 3; i++ {
		for {
			continue LABEL3
		}
	}
	fmt.Println("ok")

	for i := 0; i < 3; i++ {
	LABEL4:
		for j := 0; j < 3; j++ {
			if j > 1 {
				break LABEL4 // the same as `break`
			}
			fmt.Println(i, j)
		}
	}
}
