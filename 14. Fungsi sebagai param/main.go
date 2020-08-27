package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"Khaiz", "Badaru", "Tammam"}

	var dataContainsM = filter(data, func(each string) bool {
		return strings.Contains(each, "m")
	})

	var datalength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t \t", data)

	fmt.Println("Filter huruf M", dataContainsM) //Output sebuah kata yang mengandung huruf M

	fmt.Println("Filter Jumlah Huruf 5", datalength5) //filter sebuah kata yang terdiri dari 5 huruf
}

func filter(data []string, callback func(string) bool) []string { //fungsi ini bernilai true atau false untuk print return
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
