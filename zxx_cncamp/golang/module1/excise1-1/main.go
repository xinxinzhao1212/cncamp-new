package main

import "fmt"

func main() {

	array := [5]string{"I", "am", "stuip", "and", "weak"}

	for i, c := range array {
		fmt.Println(i, string(c))
	}
	for i := 0; i < 5; i++ {

		switch array[i] {
		case "stuip":
			array[i] = "smart"
		case "weak":
			array[i] = "strong"
		default:
		}
	}

	for i, c := range array {
		fmt.Println(i, string(c))
	}

}
