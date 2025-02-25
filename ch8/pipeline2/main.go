package main

import "fmt"

func main() {

	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()
	// 可以使用range遍历channel
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}

}
