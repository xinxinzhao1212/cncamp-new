package main

import "fmt"

func main() {
	message := make(chan int)

	//child Thread,send message
	go func() {

		message <- 1
		fmt.Println("This is from the child Thread")
	}()

	//main Thread ,receive message
	a := <-message
	fmt.Println(a)

}
