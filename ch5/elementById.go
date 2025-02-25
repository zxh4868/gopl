package main

import "golang.org/x/net/html"

func elementById(n *html.Node, id string) *html.Node {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return n
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return elementById(c, id)
	}
	return nil
}

func foreachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if pre(n) {
			return true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if foreachNode(c, pre, post) {
			return true
		}
	}
	if post != nil {
		if post(n) {
			return true
		}
	}
	return false
}

func ElementByID(n *html.Node, id string) *html.Node {
	var res *html.Node
	foreachNode(n, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					res = n
					return true
				}
			}
		}
		return false
	}, nil)

	return res
}
func main() {

}
