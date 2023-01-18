package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
	fmt.Println(re.MatchString("slimshady99"))
	fmt.Println(re.MatchString("slimsqweqhady99"))
	fmt.Println(re.MatchString("weq"))
	fmt.Println(re.MatchString("qwe"))
}
