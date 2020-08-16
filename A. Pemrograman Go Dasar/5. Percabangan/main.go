package main

import "fmt"

/*
	Yang jadi acuan dalam percabanan adalah type data bool yaitu true dan false.
	dalam Go hanya ada 2 type percabangan yaitu if else dan switch case
	Go tidak mendukung ternary operator.
*/
func main() {
	fmt.Println("///////////////////Percabangan If Else /////////////////")
	/*
		If else di go sama seperti bahasa lain, yang membedakan hanya tanda kurungnya, yang dalam Go tidak perlu ditulis
	*/

	ipk := 3.50

	if ipk >= 3.50 {
		fmt.Println("Cumlaude")
	} else if ipk >= 3.00 {
		fmt.Println("Sangat Memuaskan")
	} else if ipk >= 2.75 {
		fmt.Println("Memuaskan")
	} else if ipk >= 2.50 {
		fmt.Println("Cukup")
	} else {
		fmt.Println("Harus ikut semester pendek")
	}

	/*
		Dalam bahasa pemrograman lain, jika kondisi bloknya hanya 1 baris, bisa tidak menggunakan kurung kurawal,
		Tapi Go wajib menggunakan kurung kurawal
	*/

	fmt.Println("/////////////////// Variable Temporary pada If Else /////////////////")
	/*
		Variabel temporary adalah variabel yang hanya bisa digunakan pada blok seleksi kondisi dimana ia ditempatkan saja.
		Penggunaan variabel ini membawa beberapa manfaat, antara lain:

		- Scope atau cakupan variabel jelas, hanya bisa digunakan pada blok seleksi kondisi itu saja
		- Kode menjadi lebih rapi
		- Ketika nilai variabel tersebut didapat dari sebuah komputasi, perhitungan tidak perlu dilakukan di dalam blok
		  masing-masing kondisi.
	*/

	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	fmt.Println("/////////////////// Swith Case /////////////////")

	nilai := 90
	switch nilai {
	case 90:
		fmt.Println("Perfect")
	case 70:
		fmt.Println("Good")
	case 60:
		fmt.Println("So Bad")
	default:
		fmt.Println("Not definion")
	}

	fmt.Println("/////////////////// Swith Case Banyak Kondisi /////////////////")

	value := 3
	switch value {
	case 4, 6, 8:
		fmt.Println("Bukan Bilangan Prima")
	case 1, 2, 3, 5, 7:
		fmt.Println("Bilangan Prima")
	default:
		fmt.Println("Tidak Terdefinisi")
	}

	fmt.Println("/////////////////// Block Statement in Switch Case /////////////////")

	/*
		Tanda ini opsional, boleh dipakai boleh tidak. Bagus jika dipakai pada blok kondisi yang didalamnya ada banyak statement,
		kode akan terlihat lebih rapi dan mudah di-maintain.
	*/

	var point1 = 9

	switch point1 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}

	fmt.Println("/////////////////// Switch Case Dengan Gaya If Else /////////////////")
	var point2 = 6

	switch {
	case point2 == 8:
		fmt.Println("perfect")
	case (point2 < 8) && (point2 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println("/////////////////// Switch Case Falltrough /////////////////")
	/*
		Switch Ini tidak menhiraukan kondisi case dikarenakan jika memakai falltrough maka akan diabaikan kondisinya dan
		dilanjutkan ke case bserikutnya
	*/

	var point3 = 6

	switch {
	case point3 == 8:
		fmt.Println("perfect")
	case (point3 < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough //tetap dilanjutkan, meskipun bernilai true
	case point3 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println("/////////////////// Switch Case Bercabang /////////////////")

	var point4 = 10

	if point4 > 7 {
		switch point4 {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point4 == 5 {
			fmt.Println("not bad")
		} else if point4 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point4 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
