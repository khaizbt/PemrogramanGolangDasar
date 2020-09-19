package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var name string
	fmt.Print("type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

//dalam output error dan baris kode setelah error tidak akan dieksekusi
