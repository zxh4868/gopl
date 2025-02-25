package main

import "fmt"

var prerequests = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"discrete math":     {"intro to programming"},
	"databases":         {"data structures"},
	"distributions":     {"data structures"},
	"networks":          {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programs":          {"data structures", "computer organization"},
	"data structures": {
		"discrete math",
	},
	"formal languages": {
		"discrete math",
	},
}

func topsort(coures map[string][]string) []string {
	order := make([]string, 0, len(coures))
	visited := make(map[string]bool)
	var visitAll func(item string)
	visitAll = func(item string) {
		if !visited[item] {
			visited[item] = true
			for _, n := range coures[item] {
				visitAll(n)
			}
			order = append(order, item)
		}
	}
	for key, _ := range coures {
		if !visited[key] {
			visitAll(key)
		}
	}
	return order
}

func main() {
	for i, course := range topsort(prerequests) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
