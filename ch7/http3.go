package main

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("%.2f", d)
}

type dms map[string]dollars

func (d dms) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "%s %s\n", item, price)
	}
}

func (d dms) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := d[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %s\n", req.URL)
		return
	}
	fmt.Fprintf(w, "%s", price)
}

// func main() {

// 	db := dms{
// 		"4090":  2.08,
// 		"3090":  1.02,
// 		"4090D": 1.98,
// 		"V100":  3.28,
// 	}

// 	mux := http.NewServeMux()
// 	mux.Handle("/list", http.HandlerFunc(db.list))
// 	mux.Handle("/price", http.HandlerFunc(db.price))
// 	log.Fatal(http.ListenAndServe(":8080", mux))

// }
