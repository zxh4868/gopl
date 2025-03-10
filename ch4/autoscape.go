package main

import (
	"html/template"
	"os"
)

func main() {
	const templ = `<p>A : {{.A}} </p> <p> B : {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string
		B template.HTML
	}
	data.A = "<b>Hello</b>!"
	data.B = "<b>World</b>!"
	if err := t.Execute(os.Stdout, data); err != nil {
		panic(err)
	}

}
