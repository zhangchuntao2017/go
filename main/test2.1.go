package main

import (
	"fmt"
)

func sayHello(s string) string {
	return "hello " + s

}

func main() {
	fmt.Println(sayHello("zct"))
}
