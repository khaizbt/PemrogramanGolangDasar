package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student // nama variable adalah s1 dan student adalah struct
	s1.name = "Khaiz Badaru Tammam"
	s1.grade = 2

	fmt.Println("name : ", s1.name)
	fmt.Println("grade : ", s1.grade)
}
