package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run findlinks1.go url")
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL %s: %s\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("Error parsing HTML: %s\n", err)
		os.Exit(1)
	}
	links := visit(nil, doc)
	for _, link := range links {
		fmt.Println(link)
	}
	fmt.Println("Found", len(links), "links.")
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for m := n.FirstChild; m != nil; m = m.NextSibling {
		links = visit(links, m)
	}
	return links
}

// 深度优先遍历整个树
func visit2(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit2(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit2(links, n.NextSibling)
	}
	return links
}

// practise 5.4
func visit3(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "link" || n.Data == "script" || n.Data == "style") {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit3(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit3(links, n.NextSibling)
	}
	return links
}
