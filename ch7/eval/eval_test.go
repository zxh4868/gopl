package eval

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"text/scanner"
)

func TextEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / PI)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"f": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"f": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"f": 212}, "100"},
	}

	var prevExpr string

	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s \n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func Test_main(t *testing.T) {
	// s := "hello,世界"
	// fmt.Println(len(s))

	// for i:=0;i<len(s);i++{
	// 	fmt.Println(s[i])
	// }

	test := struct {
		expr string
		env  Env
		want string
	}{
		"4 + 3 + 8",
		Env{"f": 212},
		"100",
	}

	expr, err := Parse(test.expr)

	if err != nil {
		t.Error(err)
	}
	fmt.Println(expr)
}

func TestScan(t *testing.T) {

	var lex scanner.Scanner
	src := "x = 42 + 3.14"
	lex.Init(strings.NewReader(src))

	// 设置扫描模式
	lex.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanFloats

	// 扫描并输出标记
	for tok := lex.Scan(); tok != scanner.EOF; tok = lex.Scan() {
		fmt.Printf("Token: %s, Literal: %s\n", scanner.TokenString(tok), lex.TokenText())
	}
}
