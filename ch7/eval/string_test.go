package eval

import (
	"testing"
)

func TestString(t *testing.T) {

	tests := []struct {
		expr string
		want string
	}{
		{"-1 + -x", "(-1 + -x)"},
		{"-1 - x", "(-1 - x)"},
		{"sqrt(A / pi)", "sqrt((A / pi))"},
		{"pow(x, 3) + pow(y, 3)", "(pow(x, 3) + pow(y, 3))"},
		{"5 / 9 * (F - 32)", "((5 / 9) * (F - 32))"},
	}

	for i, tx := range tests {
		expr, err := Parse(tx.expr)
		if err != nil {
			t.Error(err)
			continue
		}

		got := expr.String()
		if got != tx.want {
			t.Fatalf("%d. got %v, expr %v", i, got, tx.want)
		}
	}
}
