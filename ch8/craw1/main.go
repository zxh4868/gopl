package main

import (
	"fmt"
	"gopl/ch5/links"
	"log"
	"os"
)

func craw(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	// 从命令行参数开始
	go func ()  {
		worklist <- os.Args[1:]	
	}()
	seen := make(map[string]bool)
	// 并发爬取链接
	for list := range worklist{
		for _,link := range list{
			if seen[link]{
				continue
			}
			seen[link] = true
			go func(link string) {
				worklist <- craw(link)
			}(link)
		}
	}
}
