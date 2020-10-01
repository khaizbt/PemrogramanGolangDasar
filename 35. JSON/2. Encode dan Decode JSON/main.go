package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Fullname string `json:"Name"`
	Age      int
}

func main() {
	fmt.Println("////Decode Array JSON ke Array Object /////")

	var jsonString = `[
		{"Name": "Khaiz Badaru Tammam", "Age": 23},
		{"Name": "Khafif Damanhuri", "Age": 16}
	]`

	var data []User

	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("User1 :", data[0].Fullname)
	fmt.Println("User2 :", data[1].Fullname)

	fmt.Println("/////Encode Object ke JSON String")

	var object = []User{{"John Wick", 27}, {"Iko Uwais", 40}}
	var jsonData, e = json.Marshal(object)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	var jsonString1 = string(jsonData)
	fmt.Println(jsonString1)

}
