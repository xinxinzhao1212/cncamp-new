package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go func() {
		println("From child thread")
		ch <- 1

	}()
	a := <-ch
	fmt.Println(a)
}
