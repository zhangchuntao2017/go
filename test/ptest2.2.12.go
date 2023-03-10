package main

import "fmt"

type Controller interface {
	SayHello()
	SayNumber(int)
	SayHi()
}

type DefaultController struct {
}

func (d DefaultController) SayHello() {
	fmt.Println("Hello world")
}

func (d DefaultController) SayNumber(number int) {
	fmt.Println(fmt.Sprintf("%d", number))
}

func (d DefaultController) SayHi() {
	fmt.Println("SayHi")
}

func main() {
	var d DefaultController
	var c Controller
	c = d

	c.SayHello()
	c.SayNumber(123)
	c.SayHi()

}
