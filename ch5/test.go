package main

import (
	"fmt"
)

type Node struct {
	Parent, FirstChild, LastChild, PrevSibling, NextSibling *Node
	Type                                                    NodeType
	Data                                                    string
	Namespace                                               string
	Attr                                                    []Attribute
}

type NodeType int
type atom struct {
	Atom string
}
type Attribute struct {
	Key, Val string
}

// Traverse recursively visits each node starting from the given root.
func Traverse(n *Node, action func(n *Node)) {
	if n == nil {
		return
	}

	// Perform the action on the current node
	action(n)

	// Recursively traverse the first child
	if n.FirstChild != nil {
		Traverse(n.FirstChild, action)
	}

	// Recursively traverse the next sibling
	if n.NextSibling != nil {
		Traverse(n.NextSibling, action)
	}
}

func main() {
	// Example tree
	root := &Node{
		Data: "root",
		FirstChild: &Node{
			Data: "child1",
			NextSibling: &Node{
				Data: "child2",
				FirstChild: &Node{
					Data: "child2.1",
				},
			},
		},
	}

	// Define an action to print the node data
	action := func(n *Node) {
		fmt.Println(n.Data)
	}

	// Traverse the tree starting from the root
	Traverse(root, action)
}
