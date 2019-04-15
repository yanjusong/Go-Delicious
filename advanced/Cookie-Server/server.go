package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// read form data to get `name`.
	name := r.PostFormValue("name")
	fmt.Println("name is:" + name)

	// read one cookie the name is `client_cookie_1` from client.
	client_cookie, _ := r.Cookie("client_cookie_1")
	fmt.Printf("%+v\n", client_cookie)

	// read completely cookies from client.
	for _, v := range r.Cookies() {
		fmt.Printf("%+v\n", v)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "server_cookie_1",
		Value:   "server_cookie_1_value",
		Expires: time.Now().Add(111 * time.Second),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "server_cookie_2",
		Value:   "server_cookie_2_value",
		Expires: time.Now().Add(111 * time.Second),
	})

	io.WriteString(w, "Hi friends, I am the server response.")
}

func main() {
	http.HandleFunc("/setcookie", setCookieHandler)
	fmt.Println("server start finished, listening at 8080")
	http.ListenAndServe(":8080", nil)
}
