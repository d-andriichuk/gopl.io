package main

import (
	"flag"
	"fmt"
	"strconv"
)

var rm = flag.Int("rm", -1, "index of removing element")

func main() {
	flag.Parse()
	stck := []int{}
	//nums := []int{}
	for _, arg := range flag.Args() {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("It's isn't a number: %v", arg)
			continue
		}
		stck = put(stck, n)
	}
	fmt.Printf("Stack:%d\n", stck)
	fmt.Printf("Stack's top:%d\n", top(stck))
	switch *rm > 0 {
	case true:
		stck = remove(stck, *rm-1)
	}
	if *rm >= 0 {
		fmt.Printf("Stack after removing at %d position:%d\n", *rm, stck)
	}
}

func top(stck []int) int {
	return stck[len(stck)-1]
}

func remove(stck []int, i int) []int {
	copy(stck[i:], stck[i+1:])
	return stck[:len(stck)-1]
}

func put(stck []int, n ...int) []int {
	// z := []int{}
	// if i+1 < cap(stck) {
	// 	z = append(z, stck[:i]...)
	// 	z = append(z, n...)
	// 	z = append(z, stck[i+1:]...)
	// }
	// return z
	return append(stck, n...)
}
