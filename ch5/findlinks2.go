package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			links, err := findLinks(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
				continue
			}
			for _, link := range links {
				fmt.Println(link)
			}
		}
	}

}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit_(nil, doc), nil
}

func visit_(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && (node.Data == "a" || node.Data == "link" || node.Data == "script" || node.Data == "style") {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}

	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit_(links, c)
	}
	return links
}
