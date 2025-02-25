package main

import "fmt"

func reverse3(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(arr []int, n int) {
	length := len(arr)
	n = n % length
	if n == 0 {
		return
	}
	reverse3(arr)
	reverse3(arr[:n])
	reverse3(arr[n:])

}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums)
	rotate(nums, 5)
	fmt.Println(nums)
	rotate(nums, 6)
	fmt.Println(nums)

}
