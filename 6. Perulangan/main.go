package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilainya adalah", i)
	}

	fmt.Println("////////////////// Perulangan Hanya Kondisi ////////////////////////")

	var y = 5
	for y < 25 {
		fmt.Println(y)
		y++
	}

	fmt.Println("////////////////// Perulangan Tanpa Kondisi ////////////////////////")
	var nilai = 0
	for {
		fmt.Println(nilai)

		nilai++
		if nilai == 50 {
			break
		}
	}

	fmt.Println("////////////////// Break and Continue ////////////////////////")

	for point := 1; point < 10; point++ { //Dilakukan perulangan mulai angka 1 hingga 10 dengan i sebagai variabel iterasi.
		//Ketika point adalah ganjil (dapat diketahui dari i % 2, jika hasilnya 1, berarti ganjil), maka akan dipaksa lanjut ke perulangan berikutnya.
		if point%2 == 1 {
			continue
		}
		//Ketika nilai lebih besar dari 8, maka perulangan akan berhenti.
		if point > 8 {
			break
		}

		fmt.Println(point)
	}

	fmt.Println("////////////////// Perulangan Bersarang ////////////////////////")

	for a := 1; a < 6; a++ {
		for j := a; j < 6; j++ {
			fmt.Print(j)
		}

		fmt.Println() //Buat pergantian line
	}

	fmt.Println("////////////////// Pemanfaatan label pada perulangan ////////////////////////")
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
