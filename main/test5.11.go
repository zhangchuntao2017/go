package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
	fmt.Println("hello world")
}
