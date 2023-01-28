package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 判断是否包含子字符串
func StringsContains(subStrings string) bool {
	return strings.Contains("aaaaa", subStrings)
}

func toString(value []byte) string {
	return string(value)
}

func toBytes(value string) []byte {
	return []byte(value)
}

func HttpByBytes() {
	url := "http://www.baidu.com"
	var body map[string]string
	body = make(map[string]string)
	body["age"] = "20"
	body["school"] = "shanghai"
	by, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(by))

	client := http.DefaultClient

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func HttpByString() {
	url := "www.baidu.com"
	request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(`{"name","zct","school","shanghai"}`))
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(content)

}

type JsonExample struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age"`
	School string `json:"university"`
}

func JsonMarshal() {
	var jex JsonExample
	jex = JsonExample{
		Name:   "Go",
		Age:    10,
		School: "Google",
	}
	by, _ := json.Marshal(jex)
	fmt.Println(by)
}

func JsonUnmarsha1() {
	var v JsonExample
	v = JsonExample{
		Name: "",
	}
}
