package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	i := 0
	lines := make([]string, 2)
	for scan.Scan() {
		lines[i] = scan.Text()
		i++
		if i == 2 {
			break
		}
	}
	if ok := isAnogramma(lines); ok {
		fmt.Fprintf(os.Stdout, "Inputed lines is an anogramma\n")
	} else {
		fmt.Fprintf(os.Stdout, "Inputed lines is not an anogramma\n")
	}
}

func isAnogramma(lines []string) bool {
	if len(lines) != 2 {
		return false
	}
	if len(lines[0]) != len(lines[1]) {
		return false
	}
	for i, j := 0, len(lines[0])-1; i < len(lines[0])-1; i, j = i+1, j-1 {
		if lines[0][i] != lines[1][j] {
			return false
		}
	}
	return true
}
