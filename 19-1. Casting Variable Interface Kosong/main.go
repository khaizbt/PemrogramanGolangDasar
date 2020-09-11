package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10 //interface kosong tidak menampilkan nilai asli, maka dari itu harus di ubah ke int
	fmt.Println(secret, "dikalikan 10 adalah", number)

	secret = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret.([]string), ", ") //casting be bentuk asli/string array
	fmt.Println(fruits, "is my faforite")

	// interface pointer
	secret = &person{"Tammam", 23}
	var name = secret.(*person).name
	fmt.Println(name)

	var person = []map[string]interface{}{
		{"name": "Tamam", "age": 22},
		{"name": "Khafif", "age": 17},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	//Bisa segala type data ditampung
	var sayur = []interface{}{
		map[string]interface{}{"name": "Kangkung", "total": 10},
		[]string{"Pepaya", "Kobis"},
		"Gudeg",
	}

	for _, each := range sayur {
		fmt.Println(each)
	}
}

type person struct {
	name string
	age  int
}
