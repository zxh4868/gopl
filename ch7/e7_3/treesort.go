package main


import "fmt"

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

func (root *Node)Sort(values []int) {
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root == nil)
}
func (root *Node)String() string{
	var vals []int
	vals = appendValues(vals, root)
	s := ""
	for _,x := range vals{
		s += fmt.Sprintf("%d ", x)
	}
	return s
}

func main(){
	vals := []int{4,5,2,1,3,7,9,10}
	var root *Node
	for _, v := range vals {
		root = add(root, v)
	}
	fmt.Println(root)


	a := []int{1,2,3,4,5,6}

	var b []int

	fmt.Printf("%p %p\n",&a, &b)

	fmt.Printf("%p ", &a[0])

	b = a
	fmt.Printf("%p %p\n",&a, &b[0])

}

