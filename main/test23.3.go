package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("You have two seconds to caludate 19 * 4")
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Times up ! The answer is 96 Did you get it")
			return
		}
	}

}
