package main

import (
	"fmt"
	"net/http"
)

func title2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("get html : %s", err)
	}
	resp.Body.Close()
	if resp.Header.Get("content-type") != "text/html" {
		return fmt.Errorf("%s has type %s, not text/html : %s", url, resp.Header.Get("content-type"))
	}
	//...输出的标题元素
	return nil
}

func main() {

}
