package main

func main() {
	message := make(chan int)

	close(message)

	//send message to a closed channel
	message <- 2

}
