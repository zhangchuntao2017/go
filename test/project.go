package main

import (
	"fmt"
	"log"
	"os"
	"testing"
	"text/template"
)

func main() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	t, err := template.ParseFiles(pwd + "chapter5/template/index.hmtl")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(t)
}

func Hello() string {
	return "Hello world"
}

func TestHello(t *testing.T) {
	result := Hello()
	want := "Hello world"
	if result == want {
		t.Logf("Hello() =  %v, want %v", result, want)
	} else {
		t.Errorf("Hello() =  %v, want %v", result, want)
	}
}

func TestHello2(excepted string) func(t *testing.T) {
	return func(t *testing.T) {
		if excepted == "Hello world" {
			t.Logf("Hello() =  %v, want %v", excepted, "Hello world")
		} else {
			t.Errorf("Hello() =  %v, want %v", excepted, "Hello world")
		}
	}
}

func TestHello3(t *testing.T) {
	t.Run("test form hello with run ", TestHello2(Hello()))
}
