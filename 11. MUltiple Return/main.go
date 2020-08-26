package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumreference = calculate(diameter)

	fmt.Printf("Luas Lingkarann \t \t : %.2f \n", area) //%f.2 maksudnya adalah jumah digit belakangnya
	fmt.Printf("Keliling Lingkaran \t %.2f \n", circumreference)

	fmt.Printf("///////////// MUltiple Return Cara Lain /////////////")

	var luas, keliling = calculateV2(diameter)
	fmt.Printf("Luas Lingkaran Menggunakan Cara 2 \t \t %.2f \n", luas)
	fmt.Printf("Keliling Lingkaran Menggunakan Cara 2 \t \t %.2f \n", keliling)
}

func calculate(d float64) (float64, float64) {
	//Hitung Luas
	area := math.Pi * math.Pow(d/2, 2) //Math.Pow digunakan untuk perpangkatan misal d/3 pangkat 2
	//Hitung Keliling
	circumference := math.Pi * d

	return area, circumference

}

func calculateV2(d float64) (area float64, keliling float64) { //Pendefinisian Fungsi diawal
	area = math.Pi * math.Pow(d/2, 2)
	keliling = math.Pi * d

	return //Udah Otomatis return yang udah didefinisiin

}
