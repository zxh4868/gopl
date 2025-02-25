package main

import (
	"fmt"
	"log"
	"time"
)

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s (%s)", msg, time.Since(start))
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func BigSlowOperation() {
	fmt.Printf("start big slow operation\n")
	time.Sleep(5 * time.Second)
	defer trace("big slow operation")()
	time.Sleep(5 * time.Second)
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d)=%d\n", x, result) }()
	result = x * x
	return result
}

func main() {
	BigSlowOperation()
	double(3)
}
