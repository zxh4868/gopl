package main

import (
	"fmt"
	"net/http"
)

type price float32

func (p price) String() string {
	return fmt.Sprintf("%.2f", p)
}

type store map[string]price

func (s store) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range s {
			fmt.Fprintf(w, "%s %s \n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := s[item]
		if ok {
			fmt.Fprintf(w, "%s\n", price)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item %q", item)
			return
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page : %s\n", req.URL)
	}

}

// func main() {
// 	s := store{
// 		"4090":  2.08,
// 		"3090":  1.02,
// 		"4090D": 1.98,
// 		"V100": 3.28,
// 	}

// 	log.Fatal(http.ListenAndServe(":8080", s))
// }
