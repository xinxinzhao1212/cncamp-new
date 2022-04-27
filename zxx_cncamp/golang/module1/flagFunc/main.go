package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	age := flag.Int("age", 18, "specify the age you set")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input name parameter is:", *name)
	fmt.Println("input parameter is", *age)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
	fullAge := fmt.Sprintf("age is %d\n", *age)
	fmt.Println(fullAge)
}
