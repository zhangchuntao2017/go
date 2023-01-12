package main

import (
	"fmt"
)

const greeting string = " hello, world "

func main() {
	greeting = "Goodbye , world"
	fmt.Println(greeting)
}
