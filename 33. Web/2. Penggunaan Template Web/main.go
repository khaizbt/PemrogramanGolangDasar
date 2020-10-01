package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var data = map[string]string{
			"Name":    "Khaiz Badaru Tammam",
			"Message": "Aku Sayang Kamu",
			"Status":  "Bekerja",
		}

		var t, err = template.ParseFiles("index.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		_ = t.Execute(writer, data)
	})

	fmt.Println("Starting development Server at http://localhost:8000")
	_ = http.ListenAndServe(":8000", nil)
}
