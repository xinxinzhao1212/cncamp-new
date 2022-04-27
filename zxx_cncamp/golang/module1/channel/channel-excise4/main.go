package main

import "fmt"

func main() {
	message := make(chan int)

	close(message)

	//Get message from a closed channel
	b := <-message
	fmt.Println(b)

}
