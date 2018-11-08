package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CountResult struct {
	Count int
	// this should be a map so the files don't get printed twice, but I'll just leave it as a note to myself.
	Files []string
}

func main() {
	counts := make(map[string]*CountResult)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, countResult := range counts {
		if countResult.Count > 1 {
			fmt.Printf("%d\t%s\t%s\n", countResult.Count, line, strings.Join(countResult.Files, " "))
		}
	}
}

func countLines(f *os.File, counts map[string]*CountResult) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = new(CountResult)
		}
		counts[input.Text()].Count++
		counts[input.Text()].Files = append(counts[input.Text()].Files, f.Name())
	}
}
