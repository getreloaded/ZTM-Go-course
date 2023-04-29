//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

//   - number of times to roll the dice
//   - number of dice used in the rolls
//   - number of sides of all the dice (6-sided, 10-sided, etc determined
//     with a variable). All dice must use the same variable for number
//     of sides, so they all have the same number of sides.
const (
	diceNumber = 2
	diceRolls  = 4
	diceSides  = 6
)

func rolltheDices() int {
	var result int = 0
	for i := 0; i < diceNumber; i++ {
		result += (rand.Intn(diceSides) + 1)
	}
	return result
}

func comment(result int) string {
	switch {

	//  - "Snake eyes": when the total roll is 2, and total dice is 2
	case result == 2 && diceNumber == 2:
		return ("Snake eyes Even")
	//  - "Even": when the total roll is even
	case result%2 == 0:
		return ("Even")
	//  - "Lucky 7": when the total roll is 7
	case result == 7:
		return ("Lucky 7 Odd")
	//  - "Odd": when the total roll is odd
	case result%2 == 1:
		return ("Odd")
	default:
		return "of no importance"
	}

}

func main() {
	sum := 0
	for i := 0; i < diceRolls; i++ {
		rtd := rolltheDices()
		sum += rtd
		fmt.Printf("dice roll number %v: %v. It is %v\n", i+1, rtd, comment(rtd))
	}
	fmt.Println("Sum of the dice rolls is", sum)

}
