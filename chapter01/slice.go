package main

import "fmt"

func main5() {
	arr := [5]int{1, 2, 3, 4, 5}
	v := arr[1:3]
	fmt.Println(v)
	fmt.Println(len(v),cap(arr))
	fmt.Println(arr[:3])
	fmt.Println(arr[:])
	fmt.Println([]int{})

	fmt.Print(append(v, 4))
}
