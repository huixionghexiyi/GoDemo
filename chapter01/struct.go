package main

type point struct {
	X int
	Y int
}
func main3() {
	d := point{1,2}
	p := &d
	println((*p).X)
	d = point{Y:3}
	println(&d)
	println((*p).X)
	println(p.Y)
}
