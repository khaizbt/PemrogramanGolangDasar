package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("Value Number A :", numberA)
	fmt.Println("Alamat Memory Number A", &numberA)

	fmt.Println("value number B :", *numberB) //karena diatas dideklarasikan numberB adalah memory numberA maka bisa memakai *
	fmt.Println("Alamat memory numberB :", numberB)

	fmt.Println("/////////// Mengubah Isi variabel /////////////")

	numberA = 5

	fmt.Println("Value Number A :", numberA)
	fmt.Println("Alamat Memory Number A", &numberA)

	fmt.Println("value number B :", *numberB)
	fmt.Println("Alamat memory numberB :", numberB) //Walaupun Value dah diubah alamat memory tetap sama
}
