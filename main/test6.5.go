package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "zct0"
	cheeses[1] = "zct1"

	var cheesesNew = make([]string, 2)
	copy(cheesesNew, cheeses)
	fmt.Println(cheesesNew)
}
