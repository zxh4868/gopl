package main

import (
	"fmt"
	"golang.org/x/net/html"
)

func soleTitle(doc html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch e := recover(); e {
		case nil:
			//没有宕机
		case bailout{}:
			err = fmt.Errorf("mutile title element")
		default:
			panic(e)
		}
	}()
	return "", err
}

func test(x int) (result int) {

	defer func() { result = 0 }()

	result = x * x
	return result
}

func main() {
	fmt.Println(test(3))
}
