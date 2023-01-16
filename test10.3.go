package main

import (
	"fmt"
)

func main() {
	name, role := "Richard", "Drummer"
	err := fmt.Errorf("The %v %v quit", role, name)
	if err != nil {
		fmt.Println(err)
	}
}
