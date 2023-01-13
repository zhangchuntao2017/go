package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "hello"
	cheeses[1] = "world"
	cheeses = append(cheeses, "zct1", "zct3", "zct6")
	fmt.Println(cheeses)
}
