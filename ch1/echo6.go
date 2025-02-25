package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", " "
	for i := 0; i < 1000; i++ {
		for _, arg := range os.Args[1:] {
			s += sep + arg
		}
	}
	fmt.Println(time.Since(start).Seconds())
	start = time.Now()
	s = ""
	for i := 0; i < 1000; i++ {
		s += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(time.Since(start).Seconds())
}
