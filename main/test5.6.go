package main

import (
	"fmt"
)

func main() {
	s := "c"
	switch s {

	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	default:
		fmt.Println("i dont know")
	}
}
