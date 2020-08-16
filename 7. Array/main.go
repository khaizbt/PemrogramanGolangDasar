package main

import "fmt"

func main() {
	var names [5]string

	names[0] = "Khaiz"
	names[1] = "Badaru"
	names[2] = "Tammam"
	names[3] = "Khafif"
	names[4] = "Damanhuri"

	//Print 1 persatu
	for _, name := range names {
		fmt.Println(name)
	}

	/*
		Variabel names dideklarasikan sebagai array string dengan alokasi elemen 4 slot.
		Cara mengisi slot elemen array bisa dilihat di kode di atas,
		yaitu dengan langsung mengakses elemen menggunakan indeks, lalu mengisinya.

		dalam Array, Jumlah index dan jumlah value dari array harus sesuai, jika jumlahnya lebih dari jumlah index maka error
		namun jika dibawah jumlah index masih bisa berjalan karena masih masuk kedalam array
	*/
	fmt.Println("////////////////////////// CARA 2 ///////////////////////////////")
	//Cara 2, Pengisian Array bisa dilakukan waktu deklarasi Variable
	var fruits = [3]string{"Gedang", "Gandul", "Salak"}

	fmt.Println("Jumlah Array adalah ", len(fruits))
	fmt.Println("Isi Array adalah", fruits)

	fmt.Println("///////////Inisiasi Array tanpa jumlah elemen//////")

	var sayur = [...]string{
		"Kangkung",
		"Tomat",
		"Kacang",
	}

	fmt.Println("Jumlah :", len(sayur))
	fmt.Println(sayur)

	fmt.Println("//////////////// Array Multi Dimensi ///////////////")
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}} //[2] adalah jumlah array,[3] adalah jumlah dalam array
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}             //Bisa juga

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	fmt.Println("///Perulangan Elemen Array menggunakan Keyword for///")

	games := [...]string{"Dota", "PUBG", "PES", "Watch Dog"}

	for i := 0; i < len(fruits); i++ {
		fmt.Println("Game :", games[i])
	}

	fmt.Println("/////////// Alokasi Array menggunakan Keyword Make")

	var frameworks = make([]string, 2)
	frameworks[0] = "Laravel"
	frameworks[1] = "Django"

	fmt.Println(frameworks)

}
