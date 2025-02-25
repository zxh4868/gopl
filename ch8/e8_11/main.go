package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func get(url string, cancelled <-chan struct{}) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Cancel = cancelled
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s", b)
}

func fetch(urls []string) (filename string, length int, err error) {
	cancelled := make(chan struct{})
	resps := make(chan string, len(os.Args[1:]))

	for _, url := range urls {
		url := url
		go func() {
			resps <- get(url, cancelled)
		}()
	}

	resp := <-resps
	close(cancelled)

	local := "base.html"
	if !strings.HasSuffix(local, ".html") {
		local += ".html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err := io.WriteString(f, resp)
	if err != nil {
		return "", 0, err
	}
	return local, n, nil
}

func main() {

}
