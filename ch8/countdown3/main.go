package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// ch := make(chan int, 1)

	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case x := <-ch:
	// 		fmt.Println(x)
	// 	case ch <- i:
	// 	}
	// }

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Pree return to abort.")
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Lanch Abort!")
			return
		}
	}
	lanch()
}

func lanch() {
	fmt.Printf("⬆***LANCH***⬆\n")
}
