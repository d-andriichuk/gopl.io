package main

import (
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	var numbers []int
	for _, arg := range os.Args[1:] {
		n := 0
		if n, err := strconv.Atoi(arg); err != nil {
			fmt.Fprintf(os.Stderr, "treesort: This is not a number - %v. Error: %v\n", n, err)
			continue
		}
		numbers = append(numbers, n)
	}
	// tree := sort(numbers)
}

func sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
