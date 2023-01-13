package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "zct1"
	cheeses[1] = "zct6"
	cheeses[2] = "zct8"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	cheeses = append(cheeses[:0], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}
