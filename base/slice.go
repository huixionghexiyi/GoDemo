package main

import "fmt"

// Slice 切片
func Slice() {
	arr := [5]int{1:1, 0:2, 2:23, 4:4}
	v := arr[1:3]
	fmt.Println(v)
	fmt.Println(len(v),cap(arr))
	fmt.Println(arr[:3])
	fmt.Println(arr[:])
	fmt.Println([]int{})

	fmt.Print(append(v, 4))
	var arr1 *[]int
	print(arr1)
}
