package main

import "fmt"

type point1 struct {
	X int
	Y int
}
func main6() {
	m := map[string]string{}
	m["name"] = "huixiong"
	m["gender"] = "male"

	m2 := map[int]point1{
		1: {1, 2},
		2: {3,4},
	}
	delete(m2,1)
	elem,ok := m2[2]
	fmt.Println(elem,ok)
}
