package main

import "fmt"

func varDeclare() {
	var number int
	var name string

	number = 100
	name = "zct"

	fmt.Println(number, name)
}

func main() {
	varDeclare()
	varDeclare2()
}

func varDeclare2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func manyVarDeclare() {
	var numberOne, numberTwo, numberThree int
	var name string
	numberOne, numberTwo, numberThree = 1, 2, 3
	name = "zct"
	fmt.Println(numberOne, numberTwo, numberThree, name)
}

func manyVarDeclareBlock() {
	var (
		numberOne, numberTwo, numberThree int
		name                              string
	)
	numberOne, numberTwo, numberThree = 1, 2, 3
	name = "zct"
	fmt.Println(numberOne, numberTwo, numberThree, name)
}
