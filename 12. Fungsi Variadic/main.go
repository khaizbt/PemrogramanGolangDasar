package main

import (
	"fmt"
	"strings"
)

func main() {
	var avg = calculate(1, 3, 7, 9, 8, 9, 4, 7)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)

	fmt.Println(msg)

	fmt.Println("//////// Fungsi Variadic Menggunakan data Slice /////////////")

	var numbers = []int{2, 5, 8, 3, 4, 1, 7}
	var average = calculate(numbers...)
	var message = fmt.Sprintf("rata-rata versi slice adalah %.2f", average)

	fmt.Println(message)

	fmt.Println("/////////// Campuran Fungsi Biasa dengan Fungsi Variadic ///////")
	hobbies("Tammam", "Mendaki", "Ngegame", "Ngoding")

	//cara 2 memakai Slice
	var myhobbies = []string{"Makan, Ngemil"}
	hobbies("Khaiz", myhobbies...)
}

func calculate(numbers ...int) float64 { //Titik titi disini digunakan untuk membebaskan jumlah parameter dari 1 variabel(variable variadic)
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers)) //Konversi int ke Float64 dahulu(casting)
	return avg
}

//Campuran Fungsi Biasa dengan Variadic

func hobbies(name string, hobbies ...string) {
	var myhobbies = strings.Join(hobbies, ", ")

	fmt.Printf("Hello My Name is \t \t %s \n", name)
	fmt.Printf("Hobby \t\t %s \n", myhobbies)
}
