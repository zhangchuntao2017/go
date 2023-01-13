package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "hello"
	cheeses[1] = "world"
	cheeses = append(cheeses, "zct")
	fmt.Println(cheeses[2])
	fmt.Println(cheeses)
}
