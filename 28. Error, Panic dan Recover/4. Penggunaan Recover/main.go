package main

/*
	Recover digunakan agar mentakover goroutine yang sedang panik
	sehingga pesan panic tidak muncul

*/
import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("Halo,", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error Occured", r)
	} else {
		fmt.Println("Aplication Running Success")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}
