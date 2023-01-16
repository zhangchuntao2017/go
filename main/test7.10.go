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
	b := &a
	b.Ice = false
	fmt.Printf("%+v\n", *b) //指针的值 *
	fmt.Printf("%+v\n", a)

	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a)

}
