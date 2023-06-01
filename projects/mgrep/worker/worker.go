package worker

import (
	"bufio"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

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

var wg sync.WaitGroup

func filesearch(DirPath string, wl chan worklist.Job) {
	dirEntry, err := os.ReadDir(DirPath)
	if err != nil {
		panic("not a directory")
	}

	wg.Add(1)
	go func() {
		for _, dirs := range dirEntry {
			if dirs.IsDir() {
				filesearch(filepath.Join(DirPath, dirs.Name()), wl)
			} else {
				var temp worklist.Job
				temp.FilePath = filepath.Join(DirPath, dirs.Name())
				wl <- temp
			}
		}
		wg.Done()
	}()
}

func Jobseeker(path string, ch chan worklist.Job) {
	filesearch(path, ch)
	go func() {
		wg.Wait()
		close(ch)
	}()

}
