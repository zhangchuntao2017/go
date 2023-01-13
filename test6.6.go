package main

import (
	"fmt"
)

func main() {
	var players = make(map[string]int)
	players["zct0"] = 21
	players["zct1"] = 22
	players["zct2"] = 21
	players["zct3"] = 21
	fmt.Println(players["zct0"])
	fmt.Println(players["zct1"])
	fmt.Println(players["zct2"])
	fmt.Println(players["zct3"])

}
