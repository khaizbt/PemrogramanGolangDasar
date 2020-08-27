package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min: //Bilangan terus di looping, agar angkanya ketemu
				min = e
			}
		}

		return min, max
	}
	var numbers = []int{10, 90, 80, 61, 89, 99}

	var min, max = getMinMax(numbers)

	fmt.Printf("Data %v \n min : %v \n max : %v", numbers, min, max)

}
