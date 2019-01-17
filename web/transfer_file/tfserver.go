package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
		http.Error(writer, "request "+fileurl+" error.", 400)
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp.Header)

	dispositionArray, ok := resp.Header["nickname"]
	disposition := ""

	if ok {
		disposition = dispositionArray[0]
	}

	if ok && len(disposition) > 0 {
		fmt.Printf("disposition is %s. size:%d\n", disposition, len(disposition))
	} else {
		fmt.Printf("can't find disposition in m.\n")
		lastIndex := strings.LastIndex(fileurl, "/")
		filename := ""
		if lastIndex < len(fileurl) {
			filename = fileurl[lastIndex+1:]
			if len(filename) == 0 {
				filename = fileurl + ".html"
			}
			if len(filename) == 0 {
				filename = "unknow.html"
			}
		}

		disposition = "attachment; filename=" + filename
	}

	writer.Header().Set("Content-Disposition", disposition)
	writer.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	io.Copy(writer, resp.Body)
}
