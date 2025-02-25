package main

import (
	"flag"
	"fmt"
	"gopl/ch5/links"
	"log"
	"strings"
)

func crawl(url string) []string {
	fmt.Println(url)
	links, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return links
}

type link struct {
	url   string
	depth int
}

func main() {

	var depth int
	flag.IntVar(&depth, "depth", 3, "-depth int")

	var items string
	flag.StringVar(&items, "urls", "", ", 分隔符")

	flag.Parse()

	init := make([]link, 0)
	for _, x := range strings.Split(items, ",") {
		init = append(init, link{x, 0})
	}

	fmt.Println(init)
	worklist := make(chan []link)
	unseenLinks := make(chan link)

	go func() {
		worklist <- init
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for url := range unseenLinks {
				foundLinks := make([]link, 0)
				for _, x := range crawl(url.url) {
					foundLinks = append(foundLinks, link{x, url.depth + 1})
				}
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)

	for x := range worklist {
		for _, link := range x {
			if !seen[link.url] && link.depth < depth {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}
