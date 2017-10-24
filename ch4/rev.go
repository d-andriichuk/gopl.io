package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := [8]int{}
	slc := a[:]
	for i, arg := range os.Args[1:] {
		if i > 7 {
			break
		}
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("This isn't a number:%v", arg)
			fmt.Fprintf(os.Stderr, "rev: %v", err)
			continue
		}
		slc[i] = n
	}
	fmt.Println("Ureversed array:", a)
	reverse(&a)
	fmt.Println("Reversed array:", a)
}

func reverse(arr *[8]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
