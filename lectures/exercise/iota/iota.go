//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	add = iota
	sub
	mul
	div
)

type Operation uint

//   - Write a receiver function that performs the mathematical operation
//     on two operands
func (op Operation) calculate(x, y int) int {

	switch op {
	case add:
		return x + y
	case sub:
		return x - y
	case mul:
		return x * y
	case div:
		return x / y
	default:
		panic("not implemented/ invalid operation")

	}
}

func main() {

	add := Operation(add)
	fmt.Println(add.calculate(2, 2)) // = 4

	sub := Operation(sub)
	fmt.Println(sub.calculate(10, 3)) // = 7)

	mul := Operation(mul)
	fmt.Println(mul.calculate(3, 3)) // = 9

	div := Operation(div)
	fmt.Println(div.calculate(100, 2)) // = 50
}
