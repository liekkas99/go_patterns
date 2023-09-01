package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Expression interface {
	Interpret() int
}

type Number struct {
	value int
}

func NewNumber(value int) *Number {
	return &Number{value: value}
}

func (n *Number) Interpret() int {
	return n.value
}

type Plus struct {
	left  Expression
	right Expression
}

func NewPlus(left, right Expression) *Plus {
	return &Plus{left: left, right: right}
}

func (p *Plus) Interpret() int {
	return p.left.Interpret() + p.right.Interpret()
}

type Minus struct {
	left  Expression
	right Expression
}

func NewMinus(left, right Expression) *Minus {
	return &Minus{left: left, right: right}
}

func (m *Minus) Interpret() int {
	return m.left.Interpret() - m.right.Interpret()
}

func parse(input string) Expression {
	tokens := strings.Split(input, " ")
	stack := []Expression{}

	for _, token := range tokens {
		if token == "+" {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			expression := NewPlus(left, right)
			stack = append(stack, expression)
		} else if token == "-" {
			right := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			expression := NewMinus(left, right)
			stack = append(stack, expression)
		} else {
			value, _ := strconv.Atoi(token)
			expression := NewNumber(value)
			stack = append(stack, expression)
		}
	}

	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return nil
}

func main() {
	input := "5 3 2 - +"
	expression := parse(input)

	result := expression.Interpret()
	fmt.Println("Result:", result)
}
