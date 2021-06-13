package main

// Pointer 指针
func Pointer() {
	var i = 1
	var p *int // 表示一个指向int的指针类型
	p = &i // 指向变量 i 的地址
	var m = &i // 将变量 i 的地址 赋值给 m
	println(&p) // 获取指针的地址
	println(&i) // 获取变量的地址
	println(*m) // 获取指针指向的值
}