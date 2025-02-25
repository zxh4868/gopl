package main

import (
	"fmt"
	"time"
)

func main()  {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n",n,fibN)	
}


func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}


func spinner(delay time.Duration) {
	for{
		// \r 表示回车，将光标移动到行首
		for _,r := range `-\|/`{
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}