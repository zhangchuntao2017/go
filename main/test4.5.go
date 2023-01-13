package main

import (
	"fmt"
)

func feedMe(portion int, eaten int) int {
	eaten += portion
	if eaten > 5 {
		fmt.Println(eaten)
		return eaten
	}
	fmt.Println(eaten)
	return feedMe(portion, eaten)
}

func main() {
	fmt.Println(feedMe(1, 0))
}
