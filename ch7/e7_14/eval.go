package e714

import (
	"fmt"
	"math"
)

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary oprator : %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.arges[0].Eval(env), c.arges[1].Eval(env))
	case "sin":
		return math.Sin(c.arges[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.arges[0].Eval(env))
	}
	panic(fmt.Sprintf("unqupported function : %q", c.fn))
}

func (t ternary) Eval(env Env) float64 {
	switch t.op1 {
	case '?':
		switch t.op2 {
		case ':':
			if t.x.Eval(env) != 0 {
				return t.y.Eval(env)
			}
			return t.z.Eval(env)
		}
		panic(fmt.Sprintf("unspported ternary operater: %q", t.op2))
	}
	panic(fmt.Sprintf("unspported ternary operater: %q", t.op1))
}
