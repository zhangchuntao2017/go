package main

import (
	"fmt"
)

func anotherFunction(a func() string) string {
	return a()
}

func main() {
	fn := func() string {
		return "function called"
	}
	fmt.Println(anotherFunction(fn))
}
