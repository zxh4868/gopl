package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)

	var name_stack []string
	var arrs []map[string]string

	fmt.Println(os.Args[1:])
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			name_stack = append(name_stack, tok.Name.Local)
			attr := make(map[string]string)
			for _, a := range tok.Attr {
				attr[a.Name.Local] = a.Value
			}
			arrs = append(arrs, attr)

		case xml.EndElement:
			name_stack = name_stack[:len(name_stack)-1]
			arrs = arrs[:len(arrs)-1]

		case xml.CharData:
			if conntainsAll(toSlice(name_stack, arrs), os.Args[1:]) {
				fmt.Printf("%s : %s ", strings.Join(name_stack, " "), tok)
			}
		}
	}
}

func toSlice(name []string, arrs []map[string]string) []string {
	var res []string
	for i, n := range name {
		res = append(res, n)

		for attr, value := range arrs[i] {
			res = append(res, attr+"="+value)
		}
	}
	return res
}

func conntainsAll(x []string, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
