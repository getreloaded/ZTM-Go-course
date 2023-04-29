//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

type Coordinates struct {
	x float64
	y float64
}

type Rectangle struct {
	vertex1 Coordinates
	vertex2 Coordinates
	vertex3 Coordinates
	vertex4 Coordinates
}

func length(rect Rectangle) float64 {
	return math.Sqrt(math.Pow((rect.vertex1.x-rect.vertex2.x), 2) + math.Pow((rect.vertex1.y-rect.vertex2.y), 2))
}
func width(rect Rectangle) float64 {
	return math.Sqrt(math.Pow((rect.vertex1.x-rect.vertex4.x), 2) + math.Pow((rect.vertex1.y-rect.vertex4.y), 2))
}

func area(rect Rectangle) float64 {
	return length(rect) * width(rect)
}

func perimeter(rect Rectangle) float64 {
	return 2 * (length(rect) + width(rect))
}

func PrintInfo(rect Rectangle) {
	fmt.Println("Area of rectangle is", area(rect))
	fmt.Println("Perimeter of rectangle is", perimeter(rect))
}

func main() {
	var rectangle Rectangle

	// assign vertices to rectangle cyclically, else the program has be coded more defensively
	// by including case where hypotenuse is calculated instead of width or length.
	rectangle.vertex1.x = 0
	rectangle.vertex1.y = 0
	rectangle.vertex2.x = 0
	rectangle.vertex2.y = 10
	rectangle.vertex3.x = 30
	rectangle.vertex3.y = 10
	rectangle.vertex4.x = 30
	rectangle.vertex4.y = 0
	PrintInfo(rectangle)

	rectangle.vertex1.x = 0
	rectangle.vertex1.y = 0
	rectangle.vertex2.x = 0
	rectangle.vertex2.y = 20
	rectangle.vertex3.x = 60
	rectangle.vertex3.y = 20
	rectangle.vertex4.x = 60
	rectangle.vertex4.y = 0
	PrintInfo(rectangle)
}
