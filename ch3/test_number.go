package main

import (
	"fmt"
	"math"
)

func main() {
	o := 0666
	fmt.Printf("%d %o %x\n", o, o, o)
	// 副词[1]表示重复使用第一个操作数，副词#表示输出相应的前缀，如0、0x、0X
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x := int64(0xdeabcdef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// 文字符号的形式是字符写在一对单引号内
	ascii := 'a'
	unicode := '鄭'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)

	for x := 0; x < 8; x++ {
		fmt.Printf("x=%d e_x = %8.3f\n", x, math.Exp(float64(x)))
	}

}
