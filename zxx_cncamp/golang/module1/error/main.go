package main

import (
	"errors"
	"fmt"
)

func main() {
	err := fmt.Errorf("This is a error!")
	fmt.Print(err)
	err1 := fmt.Errorf("This is a error1!")
	fmt.Println(err1)
	err2 := errors.New("This is a new error")
	fmt.Println(err2)

}
