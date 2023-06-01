package worklist

type Job struct {
	FilePath string
}

//  * Display the line number, file path, and complete line containing the match
type Result struct {
	FilePath   string
	LineNumber int
	Line       string
}

type Worklist struct {
	List chan Job
}

type ResultList struct {
	List chan Result
}
