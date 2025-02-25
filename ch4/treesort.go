package main

type Node struct {
	value       int
	left, right *Node
}

func add(t *Node, val int) *Node {
	if t == nil {
		return &Node{val, nil, nil}
	}
	if val < t.value {
		t.left = add(t.left, val)
	}
	if val > t.value {
		t.right = add(t.right, val)
	}
	return t
}

func appendValues(values []int, r *Node) []int {
	if r != nil {
		values = appendValues(values, r.left)
		values = append(values, r.value)
		values = appendValues(values, r.right)
	}
	return values
}

func Sort(values []int) {
	var root *Node
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func main() {

}
