package behavioral

import (
	"fmt"
	"strings"
)

// Visitor is a behavioral design pattern that lets you separate algorithms from the objects on which they operate.

// https://refactoring.guru/design-patterns/visitor

// A pattern where a component (visitor) is allowed to traverse the entire hierarchy of types.
// Implemented by propagating a single Accept() method throughout the entire hierarchy.

// This approach is often used ffor traversal
// - Alternative to Iterator
// - Hierarchy members help you traverse themselves

// 1. Visiting intrusively: violates the Open-Closed Principle
type Expression interface {
	Print(sb *strings.Builder)
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	Value float64
}

func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.Value))
}

func (d *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.visitDoubleExpression(d)
}

type AdditionExpression struct {
	Left, Right Expression
}

func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	a.Left.Print(sb)
	sb.WriteRune('+')
	a.Right.Print(sb)
	sb.WriteRune(')')
}

func (a *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.visitAdditionExpression(a)
}

// 2. Visitor that checks types: violates Open-Closed Principle when adding new expressions
func PrintExpression(e Expression, sb *strings.Builder) {
	switch r := e.(type) {
	case *DoubleExpression:
		sb.WriteString(fmt.Sprintf("%g", r.Value))
	case *AdditionExpression:
		sb.WriteRune('(')
		PrintExpression(r.Left, sb)
		sb.WriteRune('+')
		PrintExpression(r.Right, sb)
		sb.WriteRune(')')
	}
}

// 3. Using Double Dispatch
// Which function to call?
// 1. Single dispatch: depends on name of request and type of receiver
// 2. Double dispatch: depends on name of request and type of two receivers (type of visitor, type of element being visited)
type ExpressionVisitor interface {
	visitDoubleExpression(e *DoubleExpression)
	visitAdditionExpression(e *AdditionExpression)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{sb: strings.Builder{}}
}

func (ep *ExpressionPrinter) visitDoubleExpression(e *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", e.Value))
}

func (ep *ExpressionPrinter) visitAdditionExpression(e *AdditionExpression) {
	ep.sb.WriteRune('(')
	e.Left.Accept(ep)
	ep.sb.WriteRune('+')
	e.Right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

type ExpressionEvaluator struct {
	result float64
}

func NewExpressionEvaluator() *ExpressionEvaluator {
	return &ExpressionEvaluator{}
}

func (ee *ExpressionEvaluator) visitDoubleExpression(e *DoubleExpression) {
	ee.result = e.Value
}

func (ee *ExpressionEvaluator) visitAdditionExpression(e *AdditionExpression) {
	e.Left.Accept(ee)
	left := ee.result
	e.Right.Accept(ee)
	right := ee.result
	ee.result = left + right
}

func (ee *ExpressionEvaluator) Result() float64 {
	return ee.result
}
