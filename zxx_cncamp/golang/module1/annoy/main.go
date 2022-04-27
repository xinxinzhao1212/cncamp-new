package main

import "fmt"

func main() {
	resMain := add2()

	fmt.Println("resMain = %T", resMain)

	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d\n", i)
		sum := resMain(i)
		fmt.Println(sum)
	}

}

func add2() func(int) int {
	s := 0
	res := func(i int) int {
		s += i
		return s

	}
	return res
}
