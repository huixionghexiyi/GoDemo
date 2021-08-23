package main

import "fmt"

// Array 数组
func Array() {
	var a [10]int
	b := [5]int{1, 2, 3}
	c := []int{1, 2, 3}
	d := [...]int{12}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	d[0]=1
	fmt.Println(d)
}
