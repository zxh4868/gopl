package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type Reader struct {
	s string
	i int64
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func newReader(s string) *Reader {
	return &Reader{s, 0}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	r := newReader("<html><head></head><body></body></html>")
	// r.Read([]byte(""))】
	// html.Parse 接受一个实现了 io.Reader 接口的对象作为参数。这个对象应该提供要解析的HTML内容
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ouline:%v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
