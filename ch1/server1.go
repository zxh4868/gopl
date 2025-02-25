package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)                //处理器函数转换为一个符合 http.Handler 接口的类型，并将其注册到默认的 http.ServeMux 中
	log.Fatal(http.ListenAndServe(":8080", nil)) // 使用默认的多路复用器
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL PATH = %q\n", r.URL.Path)
}
