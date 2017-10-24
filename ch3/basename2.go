package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, s := range os.Args[:] {
		fmt.Fprintf(os.Stdout, "File name: %s\n", basename(s))
	}
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
