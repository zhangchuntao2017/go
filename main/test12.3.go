package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	close(messages)
	fmt.Println("Pushed two messages onto Channel with no receivers")
	time.Sleep(time.Second * 2)
	slowFunc(messages)
}
