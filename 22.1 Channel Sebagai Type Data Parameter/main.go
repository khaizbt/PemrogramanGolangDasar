package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"khaiz", "badaru", "tammam"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessage(messages)
	}

	//Eksekusi Gourutine pada IIFE
	/*
		Eksekusi goroutine tidak harus pada fungsi atau closure yang sudah terdefinisi.
		Sebuah IIFE juga bisa dijalankan sebagai goroutine baru.
		Caranya dengan langsung menambahkan keyword go pada waktu deklarasi-eksekusi IIFE-nya.
	*/

	var fruits = make(chan string)

	go func(fruit string) {
		var dataFruits = fmt.Sprintf("Nama buah %s", fruit)
		fruits <- dataFruits
	}("Mangga")

	var printFruits = <-fruits
	fmt.Println(printFruits)
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
