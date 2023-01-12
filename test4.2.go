package main

import (
	"fmt"
)

func getPrize() (int, string) {
	i := 2
	s := "goldFish"
	return i, s
}

func main() {
	a, b := getPrize()
	fmt.Printf("You won %v %v", a, b)
}
