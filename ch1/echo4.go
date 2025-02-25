package main

import (
	"fmt"
	"os"
)

// practice 1.1
func main() {
	s, sep := "", " "
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}
