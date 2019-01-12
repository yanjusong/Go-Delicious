package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := md5.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
