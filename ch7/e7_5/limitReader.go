package main

import (
	"bytes"
	"fmt"
	"io"
)


type Reader struct{
	N int64
	R io.Reader
}

func (r *Reader)Read(p []byte)(n int, err error){
	if r.N <= 0{
		return 0, io.EOF
	}
	if int64(len(p)) > r.N{
		p = p[:r.N]
	}
	n,err = r.R.Read(p)
	r.N -= int64(n)
	return
}

func limitReader(r io.Reader, n int64) io.Reader{
	return &Reader{n,r}
}

func main(){
	r := new(bytes.Buffer)
	n1,_ := r.Write([]byte("hellow long!"))
	
	fmt.Println(n1)

	lr := limitReader(r, 10)
	b := make([]byte, 100)
	n,_ := lr.Read(b)

	fmt.Println(n)



}