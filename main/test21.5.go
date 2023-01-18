package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	from, err := os.Open("./example05.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile("./example05.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer to.Close()
	a, err := io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

}
