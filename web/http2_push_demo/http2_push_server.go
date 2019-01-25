package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var httpAddr = flag.String("http", ":9090", "Listen address")

func main() {
	flag.Parse()
	http.Handle("/root/", http.StripPrefix("/root/", http.FileServer(http.Dir("root"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok {
			// Push is supported. Try pushing rather than
			// waiting for the browser request these static assets.
			if err := pusher.Push("/root/style.css", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("/root/yanjusong.png", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, indexHTML)
	})
	log.Fatal(http.ListenAndServeTLS(*httpAddr, "cert.pem", "key.pem", nil))
}

const indexHTML = `<html>
<head>
	<title>Hello World</title>
    <link rel="stylesheet" href="root/style.css">
</head>
<body>
    <h1>hello world</h1>
    <img src="root/yanjusong.png">
</body>
</html>`
