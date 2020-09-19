package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go sendData(messages)
	retriveData(messages)
}

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retriveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data"`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout, no activities uder 5 second")
			break loop

		}
	}
}
