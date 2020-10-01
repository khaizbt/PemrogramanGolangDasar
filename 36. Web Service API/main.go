package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID   string
	Name string
	Age  int
}

var data = []student{
	{"K001", "Khaiz Badaru Tammam", 22},
	{"K002", "Khafif Damanhuri", 16},
	{"B001", "Bobby Sudrajat", 45},
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Starting Web Server at localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" { //Jika methodnya POST maka data di encode ke JSON sebagai response
		var result, err = json.Marshal(data)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}
		http.Error(w, "User Not Found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
