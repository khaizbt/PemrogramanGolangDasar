package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 20
	var refrectValue = reflect.ValueOf(number)

	fmt.Println("Tipe data", refrectValue.Type())
	fmt.Println("Tipe variable", refrectValue.Int())
	fmt.Println("nilai variabel :", refrectValue.Interface())

	var s1 = &student{Name: "Khizz", Grade: 2}
	s1.getPropertyInfo()
}

//Pengaksesan Informasi Property Variable Objek

type student struct {
	Name  string
	Grade int
}

func (s student) getPropertyInfo() {
	var refrectvalue = reflect.ValueOf(s)

	if refrectvalue.Kind() == reflect.Ptr {
		refrectvalue = refrectvalue.Elem()
	}
	var reflectType = refrectvalue.Type()

	for i := 0; i < refrectvalue.NumField(); i++ {
		fmt.Println("Nama", reflectType.Field(i).Name)
		fmt.Println("tipe data", reflectType.Field(i).Type)
		fmt.Println("nilai", refrectvalue.Field(i).Interface())
		fmt.Println("")
	}
}
