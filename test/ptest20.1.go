package main

import (
	"fmt"
	"reflect"
)

func main() {
	it := reflect.TypeOf(32)
	iv := reflect.ValueOf(32)
	fmt.Printf("整数类型：%v，反射类型：%T\n,it", it, it)
	fmt.Printf("整数类型：%v，反射类型：%T\n,it", iv, iv)
}
