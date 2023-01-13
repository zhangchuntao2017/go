package main

import (
	"fmt"
)

func main() {
	var players = make(map[string]int)
	players["zct0"] = 12
	players["zct1"] = 13
	players["zct2"] = 14
	players["zct3"] = 15

	delete(players, "zct0")
	fmt.Println(players)
}
