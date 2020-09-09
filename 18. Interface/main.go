package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
}

func main() {
	var bangunDatar hitung
	bangunDatar = persegi{10.0}

	fmt.Println("========= Persegi ==========")
	fmt.Printf("luas \t \t %v \n", bangunDatar.luas())
	fmt.Printf("Keliling \t \t %v \n", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println(" ======== Lingkaran ============")
	fmt.Printf("luas \t \t %d \n", bangunDatar.luas())
	fmt.Printf("Keliling \t \t %d \n", bangunDatar.keliling())
	fmt.Printf("Keliling \t \t %d \n", bangunDatar.(lingkaran).jariJari())
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

//Luas

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2) //Kuadrat 2
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}
