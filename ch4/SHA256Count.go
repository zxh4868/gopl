package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: SHA256Count <file>")
	}
	s := os.Args[1]
	sha := sha256.Sum256([]byte(s))
	cnt := 0
	for _, x := range sha {
		cnt += int(pc[x])
	}
	fmt.Printf("%d\n", cnt)
}
