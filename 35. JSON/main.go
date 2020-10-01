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
	var jsonString = `{"Name": "Khaiz Badaru Tammam", "Age": 23}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.Fullname)
	fmt.Println("age :", data.Age)

	fmt.Println("////Decode JSON ke Map dan Interface///")

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user : ", data1["Name"])
	fmt.Println("age", data1["Age"])

	//Interface
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodedData = data2.(map[string]interface{})

	fmt.Println("nama", decodedData["Name"])
	fmt.Println("age", decodedData["Age"])

}
