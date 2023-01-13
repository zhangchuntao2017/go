package main

import (
	"fmt"
)

type Superhero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

func main() {
	e := Superhero{
		Name: "zct",
		Age:  18,
		Address: Address{
			Number: 111,
			Street: "000",
			City:   "city",
		},
	}
	fmt.Printf("%+v\n", e)
}
