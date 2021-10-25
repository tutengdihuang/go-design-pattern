package csdn

import "strings"

type Expression interface {
	Interpret(context string) bool
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (a *AndExpression) AndExpression(expr1 Expression, expr2 Expression) {
	a.expr1 = expr1
	a.expr2 = expr2
}

func (a *AndExpression) Interpret(context string) bool {
	return a.expr1.Interpret(context) && a.expr2.Interpret(context);
}

//========================
type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (o *OrExpression) OrExpression(expr1 Expression, expr2 Expression) {
	o.expr1 = expr1
	o.expr2 = expr2
}

func (o *OrExpression) Interpret(context string) bool {
	return o.expr1.Interpret(context) || o.expr2.Interpret(context);
}

//========================
type TerminalExpression struct {
	Data string
}

func (t *TerminalExpression) TerminalExpression(data string) {
	t.Data = data
}

func (t *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.Data) {
		return true
	}
	return false
}