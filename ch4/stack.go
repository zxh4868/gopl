package main

import "fmt"

var stack []int

func main() {

	var m int
	var op string
	var num int

	fmt.Scanf("%d", &m)
	for m > 0 {
		fmt.Scanf("%s", &op)
		if op == "push" {
			fmt.Scanf("%d", &num)
		}
		switch op {
		case "push":
			stack = append(stack, num)
		case "pop":
			stack = stack[:len(stack)-1]
		case "empty":
			if len(stack) == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case "query":
			fmt.Println(stack[len(stack)-1])
		}

		m--
	}
}
