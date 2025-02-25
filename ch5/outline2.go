package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElements(n *html.Node) {
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
	if n.Type == html.TextNode {
		// 去除两侧的空白字符
		trim := strings.TrimSpace(n.Data)
		if len(trim) > 0 {
			// 正确输出多行
			lines := strings.Split(trim, "\n")
			for _, line := range lines {
				fmt.Printf("%*s%s\n", depth*2, "", line)
			}
		}
	}
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s=%q ", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("/>\n", n.Data)
		} else {
			fmt.Printf(">\n", n.Data)
		}
	}
	depth++
}

func endElements(n *html.Node) {
	depth--
	if n.FirstChild != nil {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <http_url>\n", os.Args[0])
		os.Exit(1)
	}
	urls := os.Args[1:]
	for idx, url := range urls {
		if !strings.HasPrefix(url, "https") {
			urls[idx] = "https://" + url
		}
	}
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "url: %s\tfetch: %v\n", url, err)
			os.Exit(2)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "url: %s\tparse: %v\n", url, err)
		}
		forEachNode(doc, startElements, endElements)
	}

}
