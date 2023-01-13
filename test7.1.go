package main

import (
	"fmt"
)

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := Movie{
		Name:   "zct",
		Rating: 10.0,
	}
	fmt.Println(m.Name, m.Rating)
}
