package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	// 输出索引和元素
	for i, v := range a {
		fmt.Println(i, v)
	}
	//仅输出元素
	for _, v := range a {
		fmt.Println(v)
	}
	//使用数组字面量来初始化一个数组
	var q [3]int = [3]int{1, 2, 3}
	var p [3]int = [3]int{1, 2}
	fmt.Println(q[2], p[2])
	//使用省略号决定数组的长度
	r := [...]int{1, 2, 3}
	fmt.Printf("%T\n", r)
	//可以按顺序初始化，也可以给出一组索引:索引值来初始化数组
	symbol := [...]string{0: "＄", 1: "€", 2: "￥", 3: "💴", 6: "￡"}
	fmt.Printf("%d %v\n", len(symbol), symbol)

}
