package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown. Pree return to abort.")
	ticker := time.NewTicker(1 * time.Second)
	ticker.Stop()

	for i := 10; i > 0; i-- {
		select {
		case <-ticker.C:
		case <-abort:
			fmt.Println("Lanch abort!")
			return

		}
	}
	lanch()
}
func lanch() {
	fmt.Printf("⬆***LANCH***⬆\n")
}
