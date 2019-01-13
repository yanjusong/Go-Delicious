package main

import "os"
import "fmt"

func main() {

	argsWithProg := os.Args

	paramNum := len(argsWithProg)
	fmt.Printf("command line params size:%d\n", paramNum)

	fmt.Println(argsWithProg)
}
