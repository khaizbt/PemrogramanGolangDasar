package main

import "fmt"

/*
	map mirip seperti slice, hanya saja indeks yang digunakan untuk pengaksesan bisa ditentukan sendiri
	tipe-nya (indeks tersebut adalah key).

	Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. Untuk setiap data (atau value) yang disimpan,
	disiapkan juga key-nya. Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.
*/
func main() {
	var month map[string]int
	month = map[string]int{} //Default map adalah nil, maka harus di inisialisasi

	month["januari"] = 1
	month["maret"] = 3

	fmt.Println("Januari Bulan ke", month["januari"])
	fmt.Println("Maret adalah Bulan ke", month["maret"])

	fmt.Println("///////Pendefinisian Map di Awal///////")

	var province = map[string]string{"yogyakarta": "yogyakarta", "Jawa Tengah": "Semarang", "Jawa Barat": "Bandung"}
	fmt.Println(province)

	fmt.Println("///////////// Pembuatan map dengan make //////////")

	var animals = make(map[string]string)
	animals["Ayam"] = "Mamalia"
	animals["Kambing"] = "Unggas"

	fmt.Println(animals)

	fmt.Println("////////////////////Menghapus Item Map////////////")

	delete(animals, "Kambing")
	fmt.Println(animals)

	fmt.Println("////// Pendeteksi Key Map Tertentu ////////////////")

	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Key don't exist")
	}

	fmt.Println("//////////// Slice Map //////////////////")
	var gunung = []map[string]string{
		{"Yogyakarta": "Merapi", "Lokasi": "Sleman"},
		{"Jawa Tengah": "Slamet", "Lokasi": "Purwokerto"},
	}

	fmt.Println(gunung)

}
