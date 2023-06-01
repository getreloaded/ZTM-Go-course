package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randInt := generateRandInt(1, 100)

	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	finiteRandInt := generateIntFinite(3, 1, 15)
	for i := range finiteRandInt {
		fmt.Println(i)
	}
}

func generateRandInt(min, max int) <-chan int {
	output := make(chan int, 3)
	go func() {
		for {
			output <- rand.Intn(max-min+1) + min
		}
	}()

	return output
}

func generateIntFinite(count, min, max int) <-chan int {
	output := make(chan int, 1)

	go func() {

		for i := 0; i < count; i++ {
			output <- rand.Intn(max-min+1) + min
		}
		close(output)
	}()
	return output
}
