package main

import "fmt"

func main() {
	message := make(chan string)

	close(message)

	//Get message from a closed channel
	b := <-message
	fmt.Printf("Get message \"%s\" from a closed channel\n", b)

}
