package main

import "fmt"

func main() {
	data := []string{"superman", "iron man", "aquaman", "wonderwoman"}

	for _, each := range data {
		func() {
			//recover untuk IIFE dalam perulangan

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", each, "| Message:", r)
				} else {
					fmt.Println("Appliacation running succesfully")
				}
			}()

			panic("some error happen")
		}()
	}
}
