package main

import (
	"fmt"
	"gopl/ch5/links"
	"log"
	"net/url"
	"os"
	"strings"
)

func breadthFirst(f func(item string, prefix string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item, getDomain(item))...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
func getDomain(x string) string {
	parsedUrl, err := url.Parse(x)
	if err != nil {
		log.Fatal(err)
	}
	domain := parsedUrl.Hostname()
	if colonIndex := strings.LastIndex(domain, ":"); colonIndex != -1 {
		domain = domain[:colonIndex]
	}
	return domain
}
func crawlPlus(url string, prefix string) []string {
	if strings.Contains(url, prefix) {
		fmt.Println(url)
	} else {
		fmt.Println("--------------------------------")
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawlPlus, os.Args[1:])
}
