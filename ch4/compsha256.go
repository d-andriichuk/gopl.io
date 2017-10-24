package main

import (
	"crypto/sha256"
	"fmt"
	"gopl.io/ch2/popcount"
	"os"
)

func main() {
	c1 := sha256.Sum256([]byte(os.Args[1]))
	c2 := sha256.Sum256([]byte(os.Args[2]))
	difCount := 0
	for i, el := range c1 {
		difCount += 64 - popcount.PopCount(uint64(el&c2[i]))
	}
	fmt.Printf("%x\n%x\n%d\n", c1, c2, difCount)
	// n1 := 9
	// n2 := 11
	// fmt.Printf("%b\n%b\n%b\n%d\n", n1, n2, n1&n2, 64-popcount.PopCount(uint64(n1&n2)))
}
