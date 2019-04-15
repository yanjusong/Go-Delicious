package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	v := url.Values{}
	v.Set("name", "Jason")
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:8080/setcookie", body)

	// client set cookie
	cookie := http.Cookie{Name: "client_cookie_1", Value: "client_cookie_1_value", Expires: time.Now().Add(111 * time.Second)}
	req.AddCookie(&cookie)
	req.AddCookie(&http.Cookie{
		Name:    "client_cookie_2",
		Value:   "client_cookie_2_value",
		Expires: time.Now().Add(time.Duration(0) * time.Second),
	})

	// set Content-Type for body post to server normal.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	fmt.Println("\nreq content:\n", req)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("client request failed")
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("\nget response from server:\n", string(data), err)

	// read cookies from server.
	fmt.Println("\ncookies form server:")
	for _, v := range resp.Cookies() {
		fmt.Printf("%+v\n", v)
	}
}
