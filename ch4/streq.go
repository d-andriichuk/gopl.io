package main

import (
	"fmt"
	"os"
)

func main() {
	x := [...]string{"0", "1", "2", "3", "4", "5"}
	y := [...]string{"0", "1", "2", "3", "4", "5"}
	fmt.Fprintf(os.Stdout, "Is equal - %t\n", equal(x[:], y[:]))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
