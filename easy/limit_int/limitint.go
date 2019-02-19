package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

func main() {
	fmt.Println(INT_MAX)
	fmt.Println(INT_MIN)
}
