package main

import (
	"fmt"
	"os"
	//"strconv"
)

func main() {
	for _, s := range os.Args[1:] {
		fmt.Printf("Number: %s; separated number: %s\n", s, comma(s))
	}
}

func comma(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
