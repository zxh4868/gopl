package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type price float32

func (p price) String() string {
	return fmt.Sprintf("%.2f", p)
}

type store map[string]price

func (s store) list(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("./list.html")
	if err != nil {
		panic(err)
	}
	tpl.Execute(w, s)
}

func (s store) creat(w http.ResponseWriter, req *http.Request) {
	item, dollor := req.URL.Query().Get("item"), req.URL.Query().Get("price")
	_, ok := s[item]
	if ok {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "existed item:%q\n", item)
		return
	}
	p, _ := strconv.ParseFloat(dollor, 32)
	s[item] = price(p)
}
func (s store) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	_, ok := s[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	delete(s, item)
}

func (s store) update(w http.ResponseWriter, req *http.Request) {
	item, dollor := req.URL.Query().Get("item"), req.URL.Query().Get("price")
	_, ok := s[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	p, _ := strconv.ParseFloat(dollor, 32)
	s[item] = price(p)
}

func (s store) query(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := s[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s", price)
}

func main() {

	db := store{
		"4090":  2.08,
		"3090":  1.02,
		"4090D": 1.98,
		"V100":  3.28,
	}

	mux := http.NewServeMux()

	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/create", db.creat)
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe(":8080", mux))

}
