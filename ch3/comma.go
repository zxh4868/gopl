package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma1(s string) string {
	var buf bytes.Buffer
	for i, j := len(s)-1, 0; i >= 0; i-- {
		buf.WriteByte(s[i])
		j++
		if j == 3 && i != 0 {
			buf.WriteByte(',')
			j = 0
		}
	}
	var res bytes.Buffer
	for j := len(buf.Bytes()) - 1; j >= 0; j-- {
		res.WriteByte(buf.Bytes()[j])
	}
	return res.String()
}
func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	pre := n % 3
	if pre > 0 {
		buf.WriteString(s[:pre])
		buf.WriteByte(',')
	}

	for i := pre; i < n; i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < n {
			buf.WriteByte(',')
		}
	}

	return buf.String()
}

func comma_argument(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	dot := strings.IndexByte(s, '.')
	start := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		start = 1
	}
	pre := n % 3
	if dot != -1 {
		pre = (dot-start)%3 + start
		n = dot
	}
	if pre > start {
		buf.WriteString(s[start:pre])
		buf.WriteByte(',')
	}
	fmt.Println(pre)
	for i := pre; i < n; i += 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < n {
			buf.WriteByte(',')
		}
	}
	if dot != -1 {
		buf.WriteString(s[dot:])
	}
	return buf.String()

}

// 判断两个字符串是否是同文异构
func same_string(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, c := range a {
		if !strings.Contains(b, string(c)) {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(comma("123456"))
	fmt.Println(comma1("123456"))
	fmt.Println(comma2("123456"))
	fmt.Println(comma_argument("-7123456.780"))
}
