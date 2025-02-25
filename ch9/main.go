package main

import "fmt"

type Apple struct {
	PhoneName string
}

func (a Apple) Call() {
	fmt.Printf("%s有打电话功能\n", a.PhoneName)
}

func (a Apple) SendMessage() {
	fmt.Printf("%s有发信息功能\n", a.PhoneName)
}
func (a Apple) SendEmail() {
	fmt.Printf("%s有发邮件功能\n", a.PhoneName)
}

type Phone interface {
	Call()
	SendMessage()
}

func main() {
	a := Apple{PhoneName: "apple"}
	var fce Phone
	fce = a
	fmt.Println(fce)
}
