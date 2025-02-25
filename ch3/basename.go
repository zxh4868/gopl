package main

import "strings"

func basename1(s string) string {
	// 将最后一个'/'和之前的部分全部舍弃
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 保留最后一个'.'之前的所有内容
	for i := len(s); i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") //如果没有找到“/”，则slash返回-1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {

}
