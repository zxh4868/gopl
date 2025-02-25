package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// lexer 结构体封装了 Go 的标准库 scanner.Scanner, 用于扫描字符串并生成标记
type lexer struct {
	scan  scanner.Scanner // 标准库提供的扫描器
	token rune            // 当前标记的前瞻标记（lookahead token）
}

// next 使词法分析器前进到下一个标记 , Scan方法获取的是标记类型（Iden Int Float "="）
func (lex *lexer) next() { lex.token = lex.scan.Scan() }

// text 返回当前标记对应的文本
func (lex *lexer) text() string { return lex.scan.TokenText() }

// lexPanic 类型用来表示词法分析过程中发生的错误
type lexPanic string

// describle 返回一个描述当前token的字符串，用于错误报告
func (lex *lexer) describe() string {
	// 根据不同的标记类型返回相应的描述信息
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", rune(lex.token))
}

// precedence 返回给定运算符的优先级
func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

// Parse 函数时解析器的入口，它接受一个字符串形式的算术表达式作为输入，并返回一个抽象语法树（Expr）或者错误
func Parse(input string) (_ Expr, err error) {
	defer func() {
		// 捕获可能的 panic 错误，并把它们转换为返回值中的错误
		switch x := recover().(type) {
		case nil:
			// no panic
		case lexPanic:
			err = fmt.Errorf("%s", x)
		default:
			panic(x)
		}
	}()
	// 创建一个词法分析器实例，并初始化它以开始扫描字符串
	lex := new(lexer)
	lex.scan.Init(strings.NewReader(input))
	// 设置Scanner.Mode的值，使其能够扫描标识符、整数和浮点数
	lex.scan.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats
	lex.next()          // 获取第一个标记作为前瞻标记
	e := parsrExpr(lex) // 开始解析表达式
	if lex.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", lex.describe())
	}
	return e, nil
}

// 顶级函数 会调用 parseBinary 来解析二元表达式
func parsrExpr(lex *lexer) Expr {
	return parsrBinary(lex, 1)
}

// 解析二元表达式， 使用递归下降的方法，直到遇到比指定优先级低级的运算符为止
func parsrBinary(lex *lexer, prec1 int) Expr {
	// 解析左操作数（左子树）
	lhs := parsrUnary(lex)
	// fmt.Println("lhs:", lhs)

	// 循环处理当前优先级及更高优先级的运算符
	for prec := precedence(lex.token); prec >= prec1; prec-- {
		// 如果当前运算符的优先级等于当前处理的优先级
		for precedence(lex.token) == prec {
			// 获取当前处理的运算符
			op := lex.token
			lex.next() //移动到下一个token
			// 递归解析右操作数（右子树），并提高优先级
			rhs := parsrBinary(lex, prec+1)
			// 构建二元表达式
			lhs = binary{op, lhs, rhs}
		}
	}
	return lhs
}

// 解析一元表达式，比如正负号运算符
func parsrUnary(lex *lexer) Expr {
	if lex.token == '+' || lex.token == '-' {
		op := lex.token
		lex.next()
		return unary{op, parsrUnary(lex)}
	}
	return parsePrimary(lex)
}

// 解析初级表达式， 如标识符、数值字面量或括号内表达式
func parsePrimary(lex *lexer) Expr {
	switch lex.token {
	case scanner.Ident:
		id := lex.text()
		lex.next()
		if lex.token != '(' {
			return Var(id)
		}
		lex.next() // 消耗 '('
		var args []Expr
		if lex.token != ')' {
			for {
				args = append(args, parsrExpr(lex))
				if lex.token != ',' {
					break
				}
				lex.next() // 消耗 ','
			}
			if lex.token != ')' {
				msg := fmt.Sprintf("got %s, want ')'", lex.describe())
				panic(lexPanic(msg))
			}
		}
		lex.next() // 消耗')'
		return call{id, args}
	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(lex.text(), 64)
		if err != nil {
			panic(lexPanic(err.Error()))
		}
		lex.next() // 消耗数字
		return literal(f)
	case '(':
		lex.next() // 消耗 '('
		e := parsrExpr(lex)
		if lex.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", lex.describe())
			panic(lexPanic(msg))
		}
		lex.next() // 消耗')'
		return e
	}
	msg := fmt.Sprintf("unexpected %s", lex.describe())
	panic(lexPanic(msg))
}
