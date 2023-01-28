package main

import "fmt"

type erro interface {
	Error() string
}

type ErrorCode struct {
	Code    int
	Message string
}

func (e ErrorCode) Error() string {
	return fmt.Sprintf("Code : %d,Message:%s ", e.Code, e.Message)

}
func SayError() error {
	var e ErrorCode
	e.Code = 400
	e.Message = "http status code"
	return e
}

func main() {
	SayError()
}
