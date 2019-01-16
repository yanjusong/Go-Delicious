package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download(url string, localPath string, c chan bool) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(localPath)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Header)
	defer resp.Body.Close()

	// fmt.Println(string(body))
	io.Copy(f, resp.Body)
	fmt.Println("finished")

	c <- true
}

func main() {
	c := make(chan bool)
	go download("https://github.com/yanjusong/Algorithm/archive/master.zip", "D:/1.zip", c)

	<-c
}
