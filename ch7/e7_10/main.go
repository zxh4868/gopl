package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

func main() {
	val := []int{3, 2, 2, 1, 2, 3}

	fmt.Println(IsPalindrome(sort.IntSlice(val)))

}
