//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func main() {
	greeting("Naruto")
	fmt.Println(message())

	number2, number3 := returnTwoNumbers()

	fmt.Println(add(returnOneNumber(), number2, number3))
}

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greeting(personName string) {
	fmt.Println("Hola!", personName)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func message() string {
	return "Hi!"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add(number1, number2, number3 int) int {
	return number1 + number2 + number3
}

// * Write a function that returns any number
func returnOneNumber() int {
	return 42
}

// * Write a function that returns any two numbers
func returnTwoNumbers() (int, int) {
	return 10, 20
}
