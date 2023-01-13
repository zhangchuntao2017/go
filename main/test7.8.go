package main

import (
	"fmt"
	"reflect"
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

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
