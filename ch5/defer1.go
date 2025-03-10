package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer printStack()
	f(3)

}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println("================================")
	os.Stdout.Write(buf[:n])
	fmt.Println("================================")
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
