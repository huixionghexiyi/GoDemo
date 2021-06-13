package main

import "os"

// Error 错误
/** go中理解的错误是指程序的bug
*/
type MyError struct {
	msg  string
	code int
}

func (e MyError) Error() string {
	return "this is error."
}

func Error() {
	_, err := os.Open("")
	err.Error()
}
