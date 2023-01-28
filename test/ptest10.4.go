package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 66
	v := reflect.ValueOf(&x)
	// v.SetInt(55)
	vv := v.Elem()
	fmt.Printf("反射變量vv的值：%v\n", vv)
	fmt.Printf("反射變量vv能否被修改：%v", vv.CanSet())
	vv.SetInt(55)
	fmt.Printf("反射變量vv的值：%v\n", vv)
	fmt.Printf("反射變量vv能否被修改：%v", vv.CanSet())
	fmt.Printf("反射變量vv的值：%v\n", x)
}
