package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run text_content.go <url>")
		os.Exit(1)
	}
	url := os.Args[1]
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Get Method Error: %v\n", err)
		os.Exit(2)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Parse Error: %v\n", err)
		os.Exit(3)
	}
	traverse(doc)
}

func traverse(node *html.Node) {
	if node.Type == html.TextNode {
		fmt.Println(node.Data)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode && (child.Data == "script" || child.Data == "style") {
			continue
		}
		traverse(child)
	}
}
