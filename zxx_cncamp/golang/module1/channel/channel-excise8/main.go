package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Printf(" %d the child Thread!\n", i)

		}

	}()

	timer := time.NewTicker(time.Second)
	for {
		select {
		case <-ch:
			a := <-ch
			fmt.Printf("Receive data %d from channel\n", a)
		case <-timer.C:
			fmt.Println("timeout waiting from channel ch")

		}
	}

}
