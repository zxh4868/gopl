package main

import (
	"fmt"
	"gopl/ch8/craw1/links"
	"log"
	"os"
)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // 获取令牌
	list, err := links.Extract(url)
	<-tokens // 释放令牌
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // 等待发送到任务列表的数量

	n++

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist

		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
