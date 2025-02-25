package main

import (
	"fmt"
	"os"
	"unicode"
)

func remove_space(s []rune) string {
	idx := 0
	for i := 0; i < len(s); i++ {
		if i == 0 || !unicode.IsSpace(s[i]) {
			s[idx] = s[i]
			idx++
		}
		if i > 0 && unicode.IsSpace(s[i]) && !unicode.IsSpace(s[i-1]) {
			s[idx] = s[i]
			idx++
		}
	}
	return string(s[:idx])
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: remove_space <string>")
		os.Exit(1)
	}

	space_string := os.Args[1]

	fmt.Println(space_string)
	//space_string = strings.ReplaceAll(space_string, "\n", "\n")
	//fmt.Println(space_string)
	s := remove_space([]rune(space_string))
	fmt.Println(s)
}
