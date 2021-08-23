package main

import (
	"fmt"
	"os"
)

func ChannelsDemo(list []int) int{
	//c1 := make(chan int)
	//go func() {
	//	sort.Ints(list)
	//	c1 <- 1
	//}()
	//println("I`m Ok")
	//return <-c1

	requestChannels := make(chan *Request)
	quit := make(chan bool)
	go Server(requestChannels, quit)
	go Client(requestChannels, quit)
	//time.Sleep(time.Second)
	for {
		select {
		case flag := <-quit:
			if flag {
				os.Exit(1)
			}
		}
	}
}


type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(args []int) (s int) {
	for _, v := range args {
		s += v
	}
	return
}

func Client(requestChannels chan *Request, quit chan bool) {
	for i := 0; i < 2; i++ {
		request := &Request{[]int{1, 2, 3, 4, 5}, sum, make(chan int)}
		//time.Sleep(time.Second)
		requestChannels <- request
		fmt.Printf("answer: %d\n", <-request.resultChan)
	}
	quit <- true

}

func Server(requestChannels chan *Request, quit chan bool) {
	for i := 0; i < 1; i++ {
		go handle(requestChannels)
	}
	<-quit
}

func handle(requests chan *Request) {
	for req := range requests {
		req.resultChan <- req.f(req.args)
	}
}
