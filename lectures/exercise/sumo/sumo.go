// The purpose of this function is to compare two highest integers in an array.
// We leave only their difference in the array of integers at the position of the higher number.
// If both are equal then remove both from the array.
// This will be done till no integers are left or till the last integer is left in the array.
// if no integers are left then return 0 else return last integer in array

package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers = []int{100, 1, 3, 1, 8, 4, 12, 16, 50, 34}
	result := CompeteMaps(numbers)
	fmt.Println(result)
}

func CompeteMaps(numbers []int) int {
	fmt.Println(numbers)
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	var A, B, Apos, Bpos int

	// make a map from the slice to manipulate the elements easily

	numMap := make(map[int]int, len(numbers))

	for i, v := range numbers {
		numMap[i] = v
	}

	for key, value := range numMap {
		if value > B {
			B = value
			Bpos = key
		}
	}
	numMap[Bpos] = 0

	for key, value := range numMap {
		if value > A {
			A = value
			Apos = key
		}
	}

	if B == A {
		delete(numMap, Bpos)
		delete(numMap, Apos)
	} else {
		numMap[Bpos] = B - A
		delete(numMap, Apos)
	}

	// fmt.Println(numMap, len(numMap))

	var keys []int
	var newNumbers = make([]int, 0, len(numMap))

	// Make the slice of the maps.

	for k := range numMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		newNumbers = append(newNumbers, numMap[k])
	}

	return CompeteMaps(newNumbers)
}
