package main

// MyError 错误
type MyError struct {
	msg  string
	code int
}

func (e MyError) Error() string {
	return "this is error."
}

func Error() error{
	var error MyError
	return error
}
