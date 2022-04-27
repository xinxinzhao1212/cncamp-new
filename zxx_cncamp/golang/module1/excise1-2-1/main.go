package main

import (
	"fmt"
	"time"
)

func main() {

	message := make(chan int, 10)
	defer close(message)

	//Prod

	go func() {

		for i := 0; i < 10; i++ {

			message <- i
			time.Sleep(time.Second)
			fmt.Printf("This is from the child Thread %d\n", i)

		}
	}()

	//comsumer
	timer := time.NewTimer(2 * time.Second)

	for _ = range message {

		select {
		case <-message:
			msg := <-message
			fmt.Printf("receive %d from channel\n", msg)
			return

		case <-timer.C:
			fmt.Println("timeout 1 second waiting from channel ")
		}
	}
	time.Sleep(10 * time.Second)
	fmt.Println("main process exit!")
}
