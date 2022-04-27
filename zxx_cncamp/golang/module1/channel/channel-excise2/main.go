package main

import "fmt"

func main() {

	//channel 只被定义未被初始化，此时channel 为nil
	message := make(chan int)
	a := <-message
	fmt.Println(a)

}
