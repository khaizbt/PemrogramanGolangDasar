package main

import "fmt"

/*
	Defer digunakan untuk mengakhirkan eksekusi sebuah statement
	dan digunakan agar sebuh eksekusi ditempatkan di paling akhir
*/
func main() {
	defer fmt.Println("Terima Kasih") //Walaupun ditulis dipaling atas maka nanti dieksekusi di paling akhir
	fmt.Println("Selamat datang")

	orderFood("pizza")
	orderFood("nasgor")
}

//Walaupun ada return dan defer posisi ditengah tetap dieksekusi paling akhir

func orderFood(menu string) {
	defer fmt.Println("Terima kasih sudah order, Kami tunggu kunjungan anda selanjutnya")
	if menu == "pizza" {
		fmt.Println("Pilihan Tepat")
		fmt.Println("Pizza adalah makanan enak")
		return
	}

	fmt.Println("pesanan anda :", menu)
}
