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

//var maxWorkers int = 10

var args struct {
	DirPath string `arg:"positional"`
	Search  string `arg:"positional,required"`
}

func main() {
	arg.MustParse(&args)
	wl := worklist.Worklist{List: make(chan worklist.Job, 100)}
	rl := worklist.ResultList{List: make(chan worklist.Result, 100)}
	maxWorkers := 8
	path := args.DirPath
	searchTerm := args.Search

	worker.Jobseeker(path, wl.List)

	// program with only one go routine trying to access the worklist
	// workerwg.Add(1)
	// go func() {
	// 	defer workerwg.Done()
	// 	for {
	// 		job, found := <-wl.List
	// if found && len(wl.List) != 0 {
	// 	//work with the jobs till at least one is left in the worklist
	// 	worker.SearchinFile(job, &rl, searchTerm)
	// } else {
	// 	// work with the last job on the worklist
	// 	worker.SearchinFile(job, &rl, searchTerm)
	// 	close(wl.List)
	// 	break
	// }
	// 	}
	// }()

	for i := 0; i < maxWorkers; i++ {
		workerwg.Add(1)
		workerNum := i
		go func() {
			fmt.Println("   ☆   spawn worker ", workerNum)
			defer workerwg.Done()
			for {
				job, found := <-wl.List
				if !found {
					// no more jobs
					fmt.Println("   ←   exit worker ", workerNum)
					return
				}
				fmt.Printf("   ⧗   [%v] process job: %v\n", workerNum, job)
				//work with the jobs till at least one is left in the worklist
				worker.SearchinFile(job, &rl, searchTerm)
			}
		}()
	}

	resultwg.Add(1)
	go func() {

		defer resultwg.Done()
		for {
			result, found := <-rl.List
			if !found {
				return
			}
			fmt.Printf("%v: Line# %v: %v \n", result.FilePath, result.LineNumber, strings.TrimSpace(result.Line))
		}
	}()

	workerwg.Wait()
	fmt.Println("   ✓  workers all done: close results worklist")
	close(rl.List)
	resultwg.Wait()
}
