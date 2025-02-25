package e714

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5/9*(F - 32)", Env{"F": 212}, "100"},

		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},

		{"x ? y : z", Env{"x": 1, "y": 2, "z": 3}, "2"},
		{"x ? y : z", Env{"x": 0, "y": 2, "z": 3}, "3"},
	}

	var preExpr string
	for _, test := range tests {
		if test.expr != preExpr {
			fmt.Printf("\n%s\n", test.expr)
			preExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}

	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"x % 2", "unexpected '%'"},
		{"math.Pi", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{`"hello"`, "unexpected '\"'"},
		{"log(10)", `unknown function "log"`},
		{"sqrt(1, 2)", "call to sqrt has 2 args, expected 1"},

		// test ternary
		{"x ? y", "unexpected end of file"},
		{"x : y", "unexpected ':'"},
	}

	for _, test := range tests {

		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err)
		if err.Error() != test.want {
			t.Errorf("got error [%s], want [%s]", err.Error(), test.want)
		}

	}

}
