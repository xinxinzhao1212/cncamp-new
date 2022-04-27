package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 10)

	go func() {
		//假设子程序10s后才有返回值
		time.Sleep(10 * time.Second)

		ch <- 1

	}()

	timer := time.NewTimer(5 * time.Second)

	select {
	case <-ch:
		fmt.Println("receive data from child Thread!")
	case <-timer.C:
		fmt.Println("main loop do not wait anymore!")
	}

	fmt.Println("Done!")

}
