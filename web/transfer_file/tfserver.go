package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting http file sever")
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/download", handleDownload)

	err := http.ListenAndServe(":9543", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	Openfile, err := os.Open("root/index.html")
	defer Openfile.Close() //Close after function return
	if err != nil {
		//File not found, send 404
		http.Error(writer, "index.html not found.", 404)
		return
	}

	io.Copy(writer, Openfile)
}

func handleDownload(writer http.ResponseWriter, request *http.Request) {
	//First of check if Get is set in the URL
	fileurl := request.URL.Query().Get("fileurl")
	if fileurl == "" {
		//Get not set, send a 400 bad request
		http.Error(writer, "Get 'file' not specified in url.", 400)
		return
	}
	fmt.Println("Client requests: " + fileurl)

	//url := "https://github.com/yanjusong/Algorithm/archive/master.zip"
	resp, err := http.Get(fileurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Header)

	writer.Header().Set("Content-Disposition", resp.Header.Get("Content-Disposition"))
	writer.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	// writer.Header().Set("Content-Length", "416011")

	io.Copy(writer, resp.Body)
}
