package main

import (
	"fmt"
	"math"
)

/*
	Interface bisa di-embed ke interface lain
*/
func main() {
	var bangunRuang hitung = &kubus{5}

	fmt.Println("======== Kubus =========")
	fmt.Println("luas \t \t :", bangunRuang.luas())
	fmt.Println("keliling \t \t:", bangunRuang.keliling())
	fmt.Println("volume \t\t:", bangunRuang.volume())
}

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}
type hitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3) //sisi pangkat 3
}
func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6 // sisi pangkat 2 kali 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}
