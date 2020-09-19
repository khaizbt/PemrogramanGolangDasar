package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Terima Kasih")
	os.Exit(1)
	fmt.Println("selamat datang")
}

/*
	walapupun defer ditempatkan sebelum os exit akan tetap tidak diekseskusi
	karena ditengah funf=gsi program dihentikan secara paksa
*/
