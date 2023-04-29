//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

func main() {
	var favColor string = "black"
	birthYear, age := 1991, 31
	var (
		firstInitial = "K"
		lastInitial  = "B"
	)
	var ageDays int
	ageDays = age * 365

	println("My favourite color is", favColor)
	println("My initials are", firstInitial, lastInitial, "\nMy age metrics are", birthYear, age, ageDays)
}
