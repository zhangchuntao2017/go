package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	hobbies := []string{"a", "b", "c"}
	p := Person{
		Name:    "zct",
		Age:     18,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)

	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)
}
