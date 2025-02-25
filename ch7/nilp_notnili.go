package main

import (
	"io"
)

// 含有空指针的非空接口

const debug = true

// func main() {
// 	var buf *bytes.Buffer
// 	if debug {
// 		buf = new(bytes.Buffer)
// 	}
// 	f(buf)
// 	if debug {
// 		fmt.Println(string(buf.Bytes()))
// 	}
// }

func f(out io.Writer) {
	// 可能出现out的动态类型不为空，但是动态值为nil
	if out != nil {
		out.Write([]byte("hello long!!!"))
	}
}
