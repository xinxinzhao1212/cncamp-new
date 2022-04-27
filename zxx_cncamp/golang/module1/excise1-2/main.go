package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan int, 10)

	defer close(messages)

	go func() {
		// producer
		for i := 0; i < 10; i++ {
			messages <- i
			fmt.Printf("This is from child Thread %d\n", i)
			time.Sleep(1 * time.Second)
			fmt.Println("Sleep 1 second")
		}

	}()
	// consumer
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		fmt.Printf("send message: %d\n", <-messages)
	}

}
