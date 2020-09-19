package main

import "fmt"

/*
	dikarenakan defer dieksekusi diakhir blok fungsi bukan blok kondisi maka
	jika akan dieksekusi di akhir blok kondisi dapat digunakan dengan IIFE
*/
func main() {
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		func() {
			defer fmt.Println("Halo 3")
		}()
	}

	fmt.Println("halo 2") //akan tetap dieksekusi paling akhir karena memakai IIFE
}
