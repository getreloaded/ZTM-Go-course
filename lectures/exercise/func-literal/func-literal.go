//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

func report(s []string, fn func([]string)) {
	fn(s)
}

func letters(s []string) {
	value := 0
	for _, l := range s {
		for _, c := range l {
			if unicode.IsLetter(c) {
				value += 1
			}
		}
	}

	fmt.Printf("There are %d letters in the text.\n", value)
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	digit := func(s []string) {
		value := 0
		for _, l := range s {
			for _, c := range l {
				if unicode.IsDigit(c) {
					value += 1
				}
			}
		}

		fmt.Printf("There are %d digits in the text.\n", value)
	}

	report(lines, letters)
	report(lines, digit)

	report(lines, func(s []string) {
		value := 0
		for _, l := range s {
			for _, c := range l {
				if unicode.IsSpace(c) {
					value += 1
				}
			}
		}

		fmt.Printf("There are %d spaces in the text.\n", value)
	})

	punc := func(s []string) {
		value := 0
		for _, l := range s {
			for _, c := range l {
				if unicode.IsPunct(c) {
					value += 1
				}
			}
		}

		fmt.Printf("There are %d punctuations in the text.\n", value)
	}

	report(lines, punc)

}
