package main

import (
	"fmt"
	"os"
)

// practice 1.2
func main() {
	for idx, val := range os.Args {
		fmt.Println("idx:", idx, "value: ", val)
	}
}
