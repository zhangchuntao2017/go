package main

import "reflect"

// 获取切片中的最小值
func fetchNumberListMin(values []int) int {
	if len(values) < 1 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}

	var numberMin int
	numberMin = values[0]

	for _, value := range values {
		if numberMin > value {
			numberMin = value
		}
	}
	return numberMin

}

func fetchNumberMax(values []int) int {
	if len(values) < 1 {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}

	var numnberMax int
	numnberMax = values[0]

	for _, value := range values {
		if numnberMax < value {
			numnberMax = value
		}
	}
	return numnberMax
}

type Info struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Number int    `json:"number"`
}

//判断结构体Info是否存在Name属性

func hashInfoName(value string) bool {
	var info Info
	info.Name = "zct"
	info.Age = 18
	info.Number = 100

	var typeInfo reflect.Type
	typeInfo = reflect.TypeOf(info)
	if _, ok := typeInfo.FieldByName(value); ok {
		return ok
	}
	return false
}
