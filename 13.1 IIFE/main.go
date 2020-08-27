package main

import "fmt"

func main() {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}

			r = append(r, e)
		}

		return r
	}(2) //Parameter buat New Numbers

	fmt.Println("Original Number : ", numbers)
	fmt.Println("filtered number:", newNumbers)
}
