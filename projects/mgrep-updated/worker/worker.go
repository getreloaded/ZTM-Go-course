package worker

import (
	"bufio"
	"fmt"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var workerwg sync.WaitGroup

func SearchinFile(job worklist.Job, rl *worklist.ResultList, searchTerm string) {
	openfile, err := os.Open(job.FilePath)
	if err != nil {
		panic(err)
	}

	file := bufio.NewScanner(openfile)

	lineNumber := 1
	for file.Scan() {
		line := file.Text()

		if strings.Contains(line, searchTerm) {
			result := worklist.Result{
				FilePath:   job.FilePath,
				LineNumber: lineNumber,
				Line:       line,
			}
			rl.List <- result
		}
		lineNumber++
	}
}

func filesearch(DirPath string, wl chan worklist.Job) {
	dirEntry, err := os.ReadDir(DirPath)
	if err != nil {
		panic("not a directory")
	}

	wg.Add(1)
	go func() {
		fmt.Println("   ☆   spawn filesearch: ", DirPath)
		for _, dirs := range dirEntry {
			if dirs.IsDir() {
				filesearch(filepath.Join(DirPath, dirs.Name()), wl)
			} else {
				var temp worklist.Job
				temp.FilePath = filepath.Join(DirPath, dirs.Name())
				fmt.Println("   →   write job:", temp)
				wl <- temp
			}
		}
		wg.Done()
		fmt.Println("   ←   exit filesearch: ", DirPath)
	}()
}

func Jobseeker(path string, ch chan worklist.Job) {
	filesearch(path, ch)
	go func() {
		wg.Wait()
		fmt.Println("   ✕   all jobs located. close worklist")
		close(ch)
	}()
}

// func TermSearcher(searchTerm string, wl *worklist.Worklist, rl *worklist.ResultList) {
// 	workerwg.Add(1)
// 	go func() {
// 		for {
// 			defer workerwg.Done()
// 			job, found := <-wl.List
// 			if found {
// 				SearchinFile(job, rl, searchTerm)
// 			} else {
// 				break
// 			}
// 		}
// 	}()
// 	workerwg.Wait()
// }
