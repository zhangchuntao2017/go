package main

import (
	"fmt"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "zct",
		Ice:  true,
	}
	b := Drink{
		Name: "zct",
		Ice:  true,
	}

	if a == b {
		fmt.Println("a and  b are same ")
	}
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
}
