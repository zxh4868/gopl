package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {

	fmt.Printf("url \t words \t images \n")
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			continue
		}
		words, imgs, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countwordsimages: %v\n", err)
			continue
		}
		fmt.Printf("%s \t %d \t %d \n", url, words, imgs)
	}

}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return words, images, nil
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	} else if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "script" || c.Data == "style" {
			continue
		}
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return words, images
}
