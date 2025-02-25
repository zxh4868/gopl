package main

import (
	"fmt"
)

var prereqs = map[string][]string{
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

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	visitAll(keys)
	return order
}
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
