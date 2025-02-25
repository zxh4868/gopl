package main

import (
	"flag"
	"strconv"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

type Sign struct {
	x int
}

func (s *Sign) String() string {
	return strconv.FormatInt(int64(s.x), 10) + "============================"
}

func (p *Sign) Set(s string) error {
	x, err := strconv.Atoi(s)
	p.x = x
	return err
}

func SignFlag(name string, value int, usage string) *int {
	f := Sign{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.x
}

var sign = SignFlag("sign", 213, "测试")

// func main() {
// 	flag.Parse()
// 	fmt.Printf("Sleeping for %v...", *period)
// 	time.Sleep(*period)
// 	fmt.Println()

// 	fmt.Println(*sign)
// }
