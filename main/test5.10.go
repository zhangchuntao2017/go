package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("i am run after the function completes")
	fmt.Println("hello , world")
}
