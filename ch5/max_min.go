package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func max(nums ...int) int {
	if len(nums) == 0 {
		fmt.Errorf("至少需要一个整数")
		os.Exit(1)
	}
	res := nums[0]
	for _, v := range nums[1:] {
		if res > v {
			res = v
		}
	}
	return res
}
func join(sep string, strs ...string) string {
	if len(strs) == 0 {
		fmt.Errorf("至少需要一个字符串")
		os.Exit(1)
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		res += sep + strs[i]
	}
	return res
}

func elementByTagName(doc *html.Node, name ...string) (res []html.Node) {

	if doc.Type == html.ElementNode {
		for _, tag := range name {
			if doc.Data == tag {
				res = append(res, *doc)
			}
		}
	}
	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, elementByTagName(c, name...)...)
	}
	return res
}

func main() {

}
