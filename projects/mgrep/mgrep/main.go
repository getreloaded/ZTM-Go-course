package main

import (
	"fmt"
	"mgrep/worker"
	"mgrep/worklist"
	"strings"
	"sync"

	"github.com/alexflint/go-arg"
)

var workerwg sync.WaitGroup
var resultwg sync.WaitGroup

var maxWorkers int = 10

var args struct {
	DirPath string `arg:"positional"`
	Search  string `arg:"positional,required"`
}

func main() {
	arg.MustParse(&args)
	wl := worklist.Worklist{List: make(chan worklist.Job, 100)}
	rl := worklist.ResultList{List: make(chan worklist.Result, 100)}
	path := args.DirPath
	searchTerm := args.Search

	worker.Jobseeker(path, wl.List)

	// program with only one go routine trying to access the worklist

	for i := 0; i < maxWorkers; i++ {
		workerwg.Add(1)
		go func() {
			defer workerwg.Done()
			for {
				job, found := <-wl.List
				if !found {
					//work with the jobs till at least one is left in the worklist
					return
				} else {
					// work with the last job on the worklist
					worker.SearchinFile(job, &rl, searchTerm)
				}
			}
		}()
	}

	fmt.Println()
	fmt.Println()

	resultwg.Add(1)
	go func() {

		defer resultwg.Done()
		for {
			result, found := <-rl.List
			if !found {
				//print the element till their is atleast one element remaining in the list
				return
			} else {
				//print the last element and close the list
				fmt.Printf("%v: Line# %v: %v \n", result.FilePath, result.LineNumber, strings.TrimSpace(result.Line))
			}
		}
	}()

	workerwg.Wait()
	close(rl.List)
	resultwg.Wait()
}
