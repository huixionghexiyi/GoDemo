package main

import "fmt"

// go 通过结构体作为类，
type Vertex struct {
	X int
	Y int
}

// 相当于类的方法，其中v相当于 java 中的 this
func (v Vertex) Abs() int {
	return v.X
}

func Oop() {
	v := Vertex{1, 2}
	abs := v.Abs()
	fmt.Print(abs)

	var p People = Student{"huixing", "男"}
	p.eat()
	var i interface{} = "fasf"
	t,ok := i.(string)
	switch v := i.(type) {
	case string:
		fmt.Println("string",v)
	case int:
		fmt.Println("int",v)
	}
	fmt.Println(t,ok)
}

// People 接口
type People interface {
	eat()
	sleep()
}

// Student 使用struct 作为类
type Student struct {
	name   string
	gender string
}

// Student类实现eat方法，同时这个方法也是interface中定义的方法
// 这个时候可以看出Student类实现了interface接口
// 需要实现所有方法
func (s Student) eat() {
	fmt.Println("student eat.")
}

func (s Student) sleep() {
	fmt.Println("student eat.")
}
