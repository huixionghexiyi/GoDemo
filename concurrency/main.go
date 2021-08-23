package main

import "runtime"

func main() {
	result := &Vector{}
	list := Vector{1, 2, 3}
	result.DoAll(list)
	for r := range *result {
		println(r)
	}
	select {

	}
}

type Vector []float64

var numCPU = runtime.NumCPU()

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU)
	n := len(v)
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i*n/numCPU, (i+1)*n/numCPU, u, c)
	}
	for i := 0; i < numCPU; i++ {
		<-c
	}
}

func (v Vector) DoSome(i int, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u[i]
	}
	c <- 1
}
