package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go sendData(ch)
	go getData(ch)

	time.Sleep(time.Second)
	fmt.Println("main Done")
}

func sendData(ch chan<- int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
}

func getData(ch <-chan int) {
	var input int
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%d ", input)
	}
}
