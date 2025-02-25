package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Node interface {
	String() string
}

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {

	var attrs, children string

	for _, attr := range e.Attr {
		attrs += fmt.Sprintf(" %s=%q ", attr.Name.Local, attr.Value)
	}

	for _, c := range e.Children {
		children += c.String()
	}
	return fmt.Sprintf("<%s%s>%s</%s>",
		e.Type.Local, attrs, children, e.Type.Local)
}

func main() {

	node, err := parse(xml.NewDecoder(os.Stdin))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(node)
}

func parse(d *xml.Decoder) (Node, error) {
	var stack []*Element
	var parent *Element
	for {
		tok, err := d.Token()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			e := &Element{tok.Name, tok.Attr, nil}
			if len(stack) > 0 {
				parent = stack[len(stack)-1]
				parent.Children = append(parent.Children, e)
			}
			stack = append(stack, e)
		case xml.EndElement:
			if len(stack) == 1 {
				return stack[0], nil
			} else if len(stack) == 0 {
				log.Fatal("stack is empty, unexpected tag closing")
			}
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				parent = stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(tok))
			}
		}
	}
}
