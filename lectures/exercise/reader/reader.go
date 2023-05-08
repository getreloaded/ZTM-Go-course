//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	line := bufio.NewReader(os.Stdin)
	fmt.Println("---Please type something---\n commands: hello, bye")
	ComCount, LineCount := 0, 0
	for {
		line, inputErr := line.ReadString('\n')
		r := strings.TrimSpace(line)
		if r == "hello" {
			fmt.Println("Welcome to my Program.")
			ComCount++ //counts the number of commands
			continue
		}

		if r == "bye" {
			fmt.Println("See you next time.")
			ComCount++ //counts the number of commands
			continue
		}

		if r == "q" || r == "Q" {
			fmt.Println("Non-empty lines:", LineCount+ComCount) //non-command and command lines together
			fmt.Println("Command lines:", ComCount)
			break
		}

		if r != "" {
			LineCount++ // counts non command lines
		}

		if inputErr != nil {
			fmt.Println("Something is wrong:", inputErr)
		}

	}

}
