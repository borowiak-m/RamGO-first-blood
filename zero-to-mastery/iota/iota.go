package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Opeartion int

func (op Opeartion) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	}
	panic("unhandled operation")
}

func main() {
	add := Opeartion(Add)
	sub := Opeartion(Subtract)
	mul := Opeartion(Multiply)
	div := Opeartion(Divide)
	fmt.Println(add.calculate(2, 2))   // = 4
	fmt.Println(sub.calculate(10, 3))  // = 4
	fmt.Println(mul.calculate(3, 3))   // = 4
	fmt.Println(div.calculate(100, 2)) // = 4
}
