package main

import (
	"fmt"
	"time"
)

func GoRoutineDemo() {
	go spinner(100 * time.Millisecond)
	n := 45
	fibN := fib(n)
	fmt.Printf("\rfib(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
