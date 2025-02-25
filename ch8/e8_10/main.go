package main

import (
	"fmt"
	"gopl/ch8/e8_10/links"
	"log"
	"os"
)

func crawl(url string, cancled <-chan struct{}) []string {
	fmt.Println(url)
	links, err := links.Extract(url, cancled)
	if err != nil {
		log.Print(err)
	}
	return links
}

func main() {
	worklist := make(chan []string) // 可能有重复的URL列表
	unseenLinks := make(chan string)
	seen := make(map[string]bool)

	// 项任务列表中添加命令行参数
	go func() {
		worklist <- os.Args[1:]
	}()

	cancled := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancled)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link, cancled)
				// crawl(link) 是一个可能耗时的操作（例如网络请求或文件读取）。
				// 如果在 crawl(link) 完成后直接向 worklist 发送数据，而 worklist 的接收端还没有准备好，那么当前的 goroutine 会被阻塞，无法继续处理 unseenLinks 中的其他链接。
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}

	}
}
