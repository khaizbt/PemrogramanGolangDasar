package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan string)

	sayHelloTo := func(who string) {
		var data = fmt.Sprintf("Hello %s", who)
		messages <- data
	}

	go sayHelloTo("Khaiz")
	go sayHelloTo("Tammam")
	go sayHelloTo("badaru")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var messagee3 = <-messages
	fmt.Println(messagee3)
}
