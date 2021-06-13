package main

type point struct {
	X int
	Y int
}

// Struct 结构体
func Struct() {
	d := point{1,2}
	p := &d
	println((*p).X)
	d = point{Y:3}
	println(&d)
	println((*p).X)
	println(p.Y)
}
