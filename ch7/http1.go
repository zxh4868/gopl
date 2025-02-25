package main

import (
	"fmt"
	"net/http"
)

type dollor float32

func (d dollor) String() string {
	return fmt.Sprintf("%.2f", d)
}

type database map[string]dollor

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s \n", item, price)
	}
}

// func main() {
// 	db := database{
// 		"4090":  2.08,
// 		"3090":  1.02,
// 		"4090D": 1.98,
// 	}
// 	log.Fatal(http.ListenAndServe(":8080", db))
// }
