package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1, 2, 3, 4}
	sr := reflect.ValueOf(s)
	sl := sr.Len()
	fmt.Printf("Len()获取切片长度：%v，数据类型：%T\n", sl, sl)

	srp := sr.Pointer()
	fmt.Printf("Pointer()获取切片内存：%v，数据类型：%T\n", srp, srp)

	si := sr.Index(0)
	fmt.Printf("Index()读取某个元素%v，数据类型：%T\n", si, si)

	si.Set(reflect.ValueOf(666))

	s3 := sr.Slice3(0, 3, 4)

	fmt.Printf("Slice3()截取元素：%v，数据类型：%T\n", s3, s3)

	ss := sr.Slice(0, 1)
	fmt.Printf("Slice()截取元素：%v，数据类型：%T\n", ss, ss)

	srr := sr.Interface().([]int)
	fmt.Printf("反射转换切片：%v\n", srr)

	sr = reflect.Append(sr)

	fmt.Printf("Append()添加切片元素%v\n", sr)

	srrr := reflect.AppendSlice(sr, reflect.ValueOf([]int{777}))

	fmt.Printf("Append()添加切片元素%v\n", srrr)

	sss := reflect.TypeOf(s)

	fmt.Printf("Append()添加切片元素%v\n", sss)

}
