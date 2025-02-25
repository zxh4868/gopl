package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func fetch(url string) (filename string, length int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(url)
	if local == "/" {
		local = "base.html"
	}
	if !strings.HasSuffix(local, ".html") {
		local += ".html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err := io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, err
	}
	return local, n, nil
}

func main() {

	name, length, err := fetch("https://www.123dua.com/dudu-38/")
	if err == nil {
		fmt.Println(name, length)
	} else {
		fmt.Println(err)
	}

}
