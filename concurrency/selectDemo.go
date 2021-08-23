package main

import (
	"fmt"
	"time"
)

func SelectDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "channel_1"
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- "channel_2"
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Println("case 1 " + msg)
		case msg := <-ch2:
			fmt.Println("case 2 " + msg)
		}
	}

}
