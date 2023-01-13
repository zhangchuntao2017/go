package main

import （
	"fmt"
）
type Moive struct {
	Name   string
	Rating float32
}

func main() {
	m := Moive{
		Name:   "zct",
		Rating: 1.0,
	}
	fmt.Printf("%+v\n", m)
}
