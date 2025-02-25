package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

func titleForeach(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		titleForeach(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("invalid Content-Type: %s", ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		resp.Body.Close()
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	visitNode := func(n *html.Node) {
		//每个标签包裹的内容（包括文本和子标签）都被视为其子节点
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	titleForeach(doc, visitNode, nil)
	return nil
}

func main() {
	url := os.Args[1]
	if err := title(url); err != nil {
		log.Fatal(err)
	}
}
