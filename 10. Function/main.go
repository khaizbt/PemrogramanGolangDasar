package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Khaiz Badaru", "Tammam"}
	printMessage("Hello", names)

	fmt.Println("///////////////// Fungsi Dengan Return Value //////////////")

}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ") //Penggabungan antar Slice

	fmt.Println(message, nameString)
}
