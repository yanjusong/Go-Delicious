package main

import "flag"
import "fmt"

func main() {

	wordPtr := flag.String("root", "./", "a string")

	numbPtr := flag.Int("port", 9090, "an int")
	boolPtr := flag.Bool("http", true, "a bool")

	var svar string
	flag.StringVar(&svar, "cert", "./key.pem", "a string var")

	flag.Parse()

	fmt.Println("root:", *wordPtr)
	fmt.Println("port:", *numbPtr)
	fmt.Println("http:", *boolPtr)
	fmt.Println("cert:", svar)
	fmt.Println("tail:", flag.Args())
}
