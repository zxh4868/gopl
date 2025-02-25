package main

import "fmt"

// 去除[]string slice中相邻的重复字符串元素
func remove_dup(s []int) int {
	res := s[:1]
	for _, v := range s {
		if res[len(res)-1] != v {
			res = append(res, v)
		}
	}
	return len(res)
}

func main() {
	nums := []int{1, 2, 3, 4, 4, 5, 6, 5, 7, 7, 7, 7, 8, 9}

	fmt.Println(nums)
	length := remove_dup(nums)
	fmt.Println(nums[:length])
}
