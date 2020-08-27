package main

import "fmt"

func main() {
	number := 4
	fmt.Println("before", number)
	fmt.Println("after", &number)

	change(&number, 10)
	fmt.Println("after", number)
	fmt.Println("after", &number)
}

func change(original *int, value int) { //Param original dengan type pointer int
	*original = value
}
