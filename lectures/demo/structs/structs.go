package main

import "fmt"

func main() {

	newTest := []struct {
		delta     int
		testvalue int
	}{
		{200, 100},
		{-200, 0},
		{50, 50},
		{-25, 25},
	}
	for _, test := range newTest {
		fmt.Println(test.delta)
	}
}
