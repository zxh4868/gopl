package main

import "fmt"

// nonempty返回一个新的slice， slice中元素都是非空字符串
func nonempty(s []string) []string {
	i := 0
	for _, t := range s {
		if t != "" {
			s[i] = t
			i++
		}
	}
	return s[:i]
}

func nonempty2(strings []string) []string {
	s := strings[:0]
	for _, t := range strings {
		if t != "" {
			s = append(s, t)
		}
	}
	return s
}

func main() {
	strings := []string{"a", "bc", "", "cde", "defg"}
	fmt.Println(strings)
	fmt.Println(nonempty2(strings))
	fmt.Println(strings)

}
