//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode"
)

// type Count struct {
// 	count int
// 	sync.Mutex
// }

func lettercounter(word string) int {
	wordcopy := word
	total := 0
	for _, c := range wordcopy {
		if unicode.IsLetter(c) {
			total += 1
		}
	}
	return total
}

func calculate(w <-chan string, r chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	wordcopy := <-w
	r <- lettercounter(wordcopy)
}

func main() {
	textInput := bufio.NewScanner(os.Stdin)
	var wg sync.WaitGroup
	wordCh := make(chan string, 10)
	resultCh := make(chan int, 10)
	totalLetters := 0

	for textInput.Scan() {
		line := textInput.Text()
		wordlist := strings.Split(line, " ")

		for _, word := range wordlist {
			wordCh <- word
			wg.Add(1)
			go calculate(wordCh, resultCh, &wg)
		}
	}
	wg.Wait()
	//close(wordCh)
	//close(resultCh)

	for {
		r, ok := <-resultCh
		if ok {
			totalLetters += r
		} else {
			break
		}
	}
	fmt.Println(totalLetters, "letters")
}
