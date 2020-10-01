package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://facebook.com:80/hello?name=Tamam&age=27&status=single"
	var u, err = url.Parse(urlString)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("URL : %s \n", urlString)

	fmt.Printf("Protocol : %s \n", u.Scheme)
	fmt.Printf("Host : %s \n", u.Scheme)
	fmt.Printf("path : %s \n", u.Path)

	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	var status = u.Query()["status"][0]

	fmt.Printf("name : %s, age : %s, status: %s", name, age, status)
}
