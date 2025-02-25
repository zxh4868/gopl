package main

import (
	"fmt"
	"gopl/ch7/e7_16/eval"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/calc", calc)
	http.ListenAndServe("192.168.124.11:8888", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	templ := template.Must(template.ParseFiles("index.html"))
	if err := templ.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func calc(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	exprS := r.PostFormValue("expr")
	evnS := r.PostFormValue("env")
	expr, err := eval.Parse(exprS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatal(err)
	}
	env, err := parseEnv(evnS)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(expr.Eval(env))
	if err = templ.Execute(w, expr.Eval(env)); err != nil {
		log.Fatal(err)
	}
}

func parseEnv(envStr string) (eval.Env, error) {
	env := eval.Env{}

	// 使用 strings.FieldsFunc 按照指定的分隔符来分割字符串
	fileds := strings.FieldsFunc(envStr, func(r rune) bool {
		return strings.ContainsRune(`:=,{}\"`, r) || unicode.IsSpace(r)
	})

	for i := 0; i+1 < len(fileds); i += 2 {
		k := strings.TrimSpace(fileds[i])
		v := strings.TrimSpace(fileds[i+1])
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		env[eval.Var(k)] = val
	}

	return env, nil
}
