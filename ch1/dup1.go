package main

import (
	"bufio"
	"fmt"
	"os"
)

// dup1 输出标准输入中出现次数大于1的行
func main() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++
	}
	for line, num := range count {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}
	}
}
