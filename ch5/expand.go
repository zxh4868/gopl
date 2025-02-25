package main

import (
	"fmt"
	"os"
)

func expand(s string, f func(string) string) string {
	idx := 0
	var result string
	for idx < len(s) {
		if idx+4 <= len(s) && s[idx:idx+4] == "$foo" {
			result += "f(\"foo\")"
			idx += 4
		} else {
			result += s[idx : idx+1]
			idx++
		}

	}
	return result
}

func main() {

	fmt.Println("original string : ", os.Args[1])
	fmt.Println("expanded string : ", expand(os.Args[1], nil))
}
