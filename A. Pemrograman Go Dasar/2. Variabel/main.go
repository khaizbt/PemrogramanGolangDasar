package main

import "fmt"

func main() {
	//Cara 1 menggunakan Var
	var firstName string = "Khaiz Badaru"
	/*
		firstName adalah nama variabel
		string adalah type data variabel
		"Khaiz Badaru" adalah isi dari variabel
	 */

	var lastName string
	lastName = "Tammam"

	fmt.Printf("Hallo %s %s \n", firstName, lastName) //printf harus memakai \n karena tidak menghasilkan baris baru
	fmt.Println("Print dengan Println, Halo ", firstName, lastName, ", Semoga Sukses ya")
	/*
		Skema penggunaan keyword var

		var <nama-variabel> <tipe-data>
		var <nama-variabel> <tipe-data> = <nilai>

		example :
		var lastName string
		var firstName string = "john"
	 */



	fmt.Println("///////////////////////////// CARA 2 /////////////////////")
	fmt.Println("////////////////// Deklarasi var tanpa type data /////////")

	//Menggunaan var, tanoa type data, menggunakan perantara "="
	var framework = "Laravel"

	// tanpa var, tanpa type data, menggunakan perantara ":=
	bahasa := "Go Lang"

	//Tanda := hanya ada pada deklarasi awal, ex
	bahasa = "Python" //Tidak memakai := karena sudah dideklarasi,
	// BTW, deklarasi := hanya bisa digunakan dalam blok fungsi

	fmt.Println("///////////////////////////// Deklarasi Multi Variabel /////////////////////")
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	//pengisian nilai juga bisa dilakukan bersamaan saat deklarasi
	var fourth, fifth, sixth string = "empat", "lima", "enam"

	//memakai := juga bisa
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	//bisa berbeda type data
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println("///////////////////////////// Variabel Underscore /////////////////////")
	/*
		variable ini digunakan jika variabel tidak dideklarasikan, dikarenakan dalam
		Go semua variable harus digunakan, jdi istilah lain _ adalah tempat sampah untuk
		menampung nilai.

		Variabel underscore adalah predefined, jadi tidak perlu menggunakan := untuk pengisian nilai,
		cukup dengan = saja.

		Variable underscore tidak bisa digunakan, hanya untuk menampung value aja
	 */
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"


	fmt.Println("///////////////////////////// Variabel menggunakan keyword new /////////////////////")
	 /*
	 	Keyword new digunakan untuk membuat variabel pointer dengan tipe data tertentu.
	 	Nilai data default-nya akan menyesuaikan tipe datanya.
	  */

	kelas := new(string)

	fmt.Println(kelas)   // output 0x20818a220
	fmt.Println(*kelas)  // output ""

	/*
		Variabel name menampung data bertipe pointer string.
	Jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut (dalam bentuk notasi heksadesimal).
	Untuk menampilkan nilai aslinya, variabel tersebut perlu di-dereference terlebih dahulu, menggunakan tanda asterisk (*)
	 */
}
