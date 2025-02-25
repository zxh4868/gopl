package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("Commencting countdown. Press Enter to abort.")

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <- time.After(10 * time.Second):
		// 不执行任何操作
	case <- abort:
		fmt.Println("Lanch Abort!")
		return
	}
	lanch()
}

func lanch() {
	fmt.Printf("⬆***LANCH***⬆\n")
}
