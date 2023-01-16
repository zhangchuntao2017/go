package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (a *R2D2) PowerOn() error {
	if a.Broken {
		return errors.New("R2D2 is Broken ")
	} else {
		return nil
	}
	return nil
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	t := T850{
		Name: "The Terminator",
	}
	r := R2D2{
		Broken: true,
	}

	err := Boot(&r)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on !")
	}

	err = Boot(&t)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on !")
	}

}
