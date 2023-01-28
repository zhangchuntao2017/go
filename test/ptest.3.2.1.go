package main

import (
	"fmt"
	"reflect"
)

var add = func(numberOne int, numberTwo int) int {
	return numberOne + numberTwo
}

var minus = func(numberOne int, numberTwo int) int {
	return numberOne - numberTwo
}

var multiply = func(numberOne int, numberTwo int) int {
	return numberOne * numberTwo
}

var division = func(numberOne float64, numberTwo int) float64 {
	return float64(float64(numberOne) / float64(numberTwo))
}

var judge = func(manAge int, womanAage int) bool {
	if manAge > 22 && womanAage >= 20 {
		return true
	}
	return false
}

var opList = func(number [4]int) {
	fmt.Println(number[1], reflect.TypeOf(number[1]))
	fmt.Println(len(number))
	fmt.Println(number[1:], reflect.TypeOf(number[1:]))

	for index, value := range number {
		fmt.Println(index, value)
	}

	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
}

var opSlice = func(name []string) []string {
	fmt.Println(name[1], reflect.TypeOf(name))

	for index, one := range name {
		fmt.Println(index, one)
	}

	name = append(name, "zct")
	return name
}

var opMap = func(name map[string]int) {
	for key, value := range name {
		fmt.Println(key, value)
	}
	name["zct"] = 100

	if value, ok := name["GO"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("not  exist GO")
	}

	delete(name, "java")

}
