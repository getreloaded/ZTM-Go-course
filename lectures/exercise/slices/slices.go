//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

// * Create a function to print out the contents of the assembly line
func PrintPart(parts []Part) {

	for p := 0; p < len(parts); p++ {
		j := p
		fmt.Printf("%v ", parts[j])
	}
	fmt.Println()
}

func main() {
	//  - Create an assembly line having any three parts
	assemblyLine := make([]Part, 3, 5)

	assemblyLine = []Part{"screw", "bolt", "nut"}
	PrintPart(assemblyLine)
	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, []Part{"allenkey", "stud"}...)
	PrintPart(assemblyLine)
	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	PrintPart(assemblyLine)

}
