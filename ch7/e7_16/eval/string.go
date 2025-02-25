package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	var arges []string

	for _, x := range c.arges {
		arges = append(arges, x.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(arges, ", "))
}
