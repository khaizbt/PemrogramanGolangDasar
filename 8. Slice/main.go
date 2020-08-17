package main

import "fmt"

/*
	Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
	Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain
	yang memiliki alamat memori yang sama.
*/
func main() {
	//Cara pembuatan Slice sama sperti pembuatan Array, hanya saja di Slice tidak ditulis jumlah Indexnya

	fruits := []string{"Banana", "Lemon", "Durians", "Orange", "Mango"}

	fmt.Println(fruits)

	/*
		Array dan Slice tidak dapat dibedakan karena memiliki kesatuan yang berbeda. Array adalah kumpulan nilai atau elemen,
		sedang slice adalah referensi tiap elemen tersebut. lice bisa dibentuk dari array yang sudah didefinisikan,
		caranya dengan memanfaatkan teknik 2 index untuk mengambil elemen-nya. Contoh sebagai berikut
	*/
	aFruits := fruits[0:2]
	fmt.Println(aFruits) //Output Index 0 sampapi sebelum 2(index ke 1)
	bfruits := aFruits[0:1]
	fmt.Println(bfruits) //Slice bisa dicustom kembali walaupun itu hasil dari Slice 2 Index
	//Hasilnya adalah banana
	//Teknik 2 Index juga bisa mengubah element dari banana ke Gedang

	bfruits[0] = "Gedang"
	fmt.Println(bfruits)
	fmt.Println(fruits)

	fmt.Println("////////////////////// Fungsi Cap() dan Len() pada Slice ////////////////////////")
	/*
		Fungsi Len digunakan untuk menghitung jumlah index yang dideklarasikan
		Fungsi Cap digunakan untuk menghitung lebar dari suatu index atau kapasitas maxsimum slice
	*/

	fmt.Println(len(fruits)) //Element dihitung semua
	fmt.Println(cap(fruits)) //Element dihitung semua
	fmt.Println("--------------------------------")
	fmt.Println(len(fruits[0:3])) //Element dihitung dari index 0 sampe index sebelum 3
	fmt.Println(cap(fruits[0:3])) //Element dihitung dari index 0 sampai index maksimum slice tersebut
	fmt.Println("--------------------------------")
	fmt.Println(len(fruits[1:3])) //dihitung dari index 1 sampai index sebelum 3
	fmt.Println(cap(fruits[1:3])) //dihitung dari index 1 sampai index terakhir dalam slice tersebut

	fmt.Println("///////////////////// Fungsi Append //////////////////////////")

	var cities = []string{"Yogyakarta", "Solo", "Semarang", "Purwokerto"}

	fmt.Println(cities[0:3])

	citiesA := cities[0:4]
	citiesA = append(cities, "Cilacap") //Maka pada variabel citiesA akan ditambahkan Cilacap

	fmt.Println(citiesA)
	fmt.Println(cities)

	fmt.Println("//////////////////////// Fungsi Copy ///////////////////////////")

	aseans := make([]string, 3) //Isi slice adalah 3
	nations := []string{"Indonesia", "Thailand", "Malaysia", "Vietnam"}

	copies := copy(aseans, nations) //Asean mencopy/mengambil element dari nations dengan jumlah 3

	fmt.Println(nations)
	fmt.Println(aseans)
	fmt.Println(copies) //Jumlah Yang dicopy

	dst := []string{"potato", "hiu", "potato"} //index pertama akan ditimpa oleh watermelon
	src := []string{"watermelon"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple potato
	fmt.Println(src) // watermelon pinnaple
	fmt.Println(n)

	fmt.Println("////////////////////// Pengaksesan Slice dengan 3 Index //////////////////////")
	var buah = []string{"apple", "grape", "banana"}
	var abuah = buah[0:2]
	var bBuah = buah[0:2:3] //3 disini bisa diisi Maksimal dari jumlah slice itu sendiri

	fmt.Println(abuah)
	fmt.Println(bBuah)
	fmt.Println(cap(bBuah))

}
