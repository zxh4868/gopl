package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // 使用栈来存储元素
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
			case xml.StartElement:
				stack = append(stack, tok.Name.Local) // 入栈
			case xml.EndElement:
				stack = stack[:len(stack)-1] // 出栈
			case xml.CharData:
				if containsAll(stack, os.Args[1:]) {
					fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
				}
		}
	}
}

// containsAll 判断x是否包含y中的所有元素
func containsAll(x, y []string) bool {
	for len(y) <= len(x){
		if len(y) == 0{
			return true
		}
		if x[0] == y[0]{
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
