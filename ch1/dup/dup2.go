package main

import (
	"bufio"
	"fmt"
	"os"
)

type Duplicates map[string]int

func main() {
	//counts := make(map[string]int)
	dupFiles := make(map[string]Duplicates)
	files := os.Args[1:]
	if len(files) == 0 {
		//countLines(os.Stdin, counts)
		countLinesWithFileName(os.Stdin, dupFiles, "terminal:")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			//countLines(f, counts)
			countLinesWithFileName(f, dupFiles, arg)
			f.Close()
		}
	}
	for name, file := range dupFiles {
		//for line, n := range counts {
		for line, n := range file {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", name, n, line)
			}
		}
	}
}

//func countLines(f *os.File, counts map[string]int) {
func countLines(f *os.File, counts Duplicates) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "exit" {
			break
		}
	}
}

func countLinesWithFileName(f *os.File, counts map[string]Duplicates, fileName string) {
	input := bufio.NewScanner(f)
	c := make(map[string]int)

	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		c[input.Text()]++
	}
	counts[fileName] = c
}
