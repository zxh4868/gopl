package main

import (
	"fmt"
	"unicode/utf8"
)

// HasPrefix 判断某个字符串是否是另一个的前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix 判断某一个字符串是否是另一个的后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func main() {
	s := "hello zxh!"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[0:5])
	s = "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", size, r)
		i += size
	}
	fmt.Println(string(19990))
	fmt.Println(string(1234567))

}
