package main

import (
	"fmt"
	"time"
)

type Filter func(c *Context)

type FilterBuilder func(next Filter) Filter

// MetricFilterBuilder 使用闭包特性 执行传入的方法，并且在执行前后计算实践
func MetricFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		startTime := time.Now().Nanosecond()
		next(c)
		endTime := time.Now().Nanosecond()
		fmt.Printf("run time: %d \n", endTime-startTime)
	}
}
