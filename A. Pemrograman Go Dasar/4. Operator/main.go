package main

import "fmt"

func main() {
	//Operator Aritmatika
	value := 3 + 3*6/(9+1)%2
	/*
		Go Mendukung penjumlahan, Pengurangan, Perkalian, Pembagian, dan Modulus(Sisa Bagi
	*/
	fmt.Println(value)

	fmt.Println("///////////////////////Operator Perbandingan /////////////////////")
	var key = (((2+6)%3)*4 - 2) / 3
	var isEqual = (key == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	fmt.Println(`//////////////////// Operator Logika \\\\\\\\\\\\\\\\\\\\\\\\`)
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)

	/*
		leftAndRight bernilai false, karena hasil dari false dan true adalah false.
		leftOrRight bernilai true, karena hasil dari false atau true adalah true.
		leftReverse bernilai true, karena negasi (atau lawan dari) false adalah true.
	*/
}
