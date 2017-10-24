package main

import (
	"fmt"
	"os"
)

func main() {
	for _, s := range os.Args[:] {
		fmt.Fprintf(os.Stdout, "File mame: %s\n", basename(s))
	}
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
