package main

import "fmt"

type Info struct {
	Name string
	_    int
	Aage int
}

func main() {
	var infoOne Info = Info{
		Name: "zct",
		Aage: 20,
	}
	var infoTwo = Info{
		"xiexie",
		200,
		2,
	}
	var infoThree = new(Info)
	infoThree = &Info{
		Name: "zct",
		Aage: 20,
	}
	fmt.Println(infoOne, infoTwo, infoThree)

}
