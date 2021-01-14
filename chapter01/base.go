package main

import (
	"fmt"
	"os"
)

// 常量
const Pi = 3.14159265354
const a string = "aaa"
const b = "bbb"
const c1 = 2 / 3
const c2 = 0.634123412541241241241246
const log = 1 / c2
const (
	RED Color = iota
	ORANGE
	GREEN
	BLUE
	INDIGO
)

type Color int

// 变量
var i1 int
var i2 bool
var i3 string
var (
	i4 = 15
	i5 = false
	i6 = "Go says hello"
	i7 = 50
)
var i8 int64 = 2
var i9 = 2
var (
	i10 = os.Getenv("HOME")
	i11 = os.Getenv("USER")
	i12 = os.Getenv("GOROOT")
)

type Options struct {
	a int
	b int
}

func f(x int) {
	a := 1
	a1 := &a // 取地址
	a2 := &a
	_, a3 := a1, a2
	fmt.Println("x:", x)
	fmt.Println("a:", a)
	fmt.Println("&a:", &a)
	fmt.Println("*a1:", *a1)
	fmt.Println(a2 == a1)
	fmt.Println(a1 == a3)
}

// init 方法，会在main方法之前执行。
func init() {
	// init
	i1 = 12
}

// main 方法
func main1() {
	//f(1)
	//a := uint(10)
	//fmt.Println(a)
	//
	//c := complex64(31 + 1i)
	//println(c)
	Multiply(1, 2, 3, 4, 54, 65, 6)
}

// 多参数函数
func Multiply(c, b int, a ...int) int {
	if len(a) == 0 {
		fmt.Println(0)
	}
	mul := a[0]
	for m, v := range a {
		fmt.Println(m)
		fmt.Println(v)
	}
	return mul
}

// 没有泛型，使用 空接口
func typeCheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Println(v)
		case float32:
			fmt.Println(v)
		}
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}
