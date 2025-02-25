package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

// practise 5.2
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run element_conut.go url1 <url2> <urln>")
		os.Exit(1)
	}
	urls := os.Args[1:]
	for i := 0; i < len(urls); i++ {
		if !strings.HasPrefix(urls[i], "https") {
			urls[i] = "https://" + urls[i]
		}
	}
	for _, url := range urls {
		counts := make(map[string]int)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Http Get Error:", err)
			os.Exit(2)
		}
		defer resp.Body.Close()
		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println("Parse Error:", err)
			os.Exit(3)
		}
		label_count(counts, doc)
		fmt.Printf("=========================url : %s\n", url)
		for k, v := range counts {
			fmt.Printf("%s: %d\n", k, v)
		}
	}

}
func label_count(counts map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counts[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		label_count(counts, c)
	}
}
