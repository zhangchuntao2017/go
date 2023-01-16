package main

import (
	"fmt"
)

func main() {
	s := "After a backslash , certain single character escapes represent special values\n\t n is a  line feed or new line \n\t t is a tab"
	fmt.Println(s)
}
