package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Commencting countdown!")

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	lanch()
}

func lanch() {

	fmt.Printf("⬆***LANCH***⬆\n")
}
