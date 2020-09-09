package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{"Khaiz Badaru", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("Nama Panggilan: ", name)
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}
