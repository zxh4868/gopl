package e714

type Expr interface {
	// Eval 返回表达式在env上下文的值
	Eval(env Env) float64
	//
	Check(vars map[Var]bool) error
	
}

type Var string

type literal float64

// unary 表示一元运算符表达式
type unary struct {
	op rune
	x  Expr
}

// binary 表示二元运算符表达式
type binary struct {
	op   rune
	x, y Expr
}

// call 表示一个函数表达式
type call struct {
	fn    string // "pow"、"sin"、"sqrt"之中的一个
	arges []Expr
}


type ternary struct{
	op1, op2 rune
	x,y,z Expr
}