package main

import (
	"bufio"
	"fmt"
	"gopl/ch7/eval"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	fmt.Println("请输入一个算术表达式")
	var exps string
	// fmt.Fscanln(reader, &exps)
	exps, _ = reader.ReadString('\n')

	expr, err := eval.Parse(exps)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return
	}

	vars := map[eval.Var]bool{}

	err = expr.Check(vars)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return
	}
	env := eval.Env{}
	for v := range vars {
		for {
			fmt.Printf("请输入 %s 的值: ", v)
			input, _ := reader.ReadString('\n') // 读取整行输入
			input = strings.TrimSpace(input)    // 去除换行符和空格

			// 尝试将输入转换为浮点数
			x, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("ERROR: 输入值无效，请输入一个数字")
				continue // 如果输入无效，提示用户重新输入
			}
			env[v] = x
			break // 输入有效，退出循环
		}
	}
	fmt.Printf("%s.Eval() in %v = %g\n", expr, env, expr.Eval(env))
}
