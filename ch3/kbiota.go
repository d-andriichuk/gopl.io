package main

import (
	"fmt"
	"os"
)

const (
	B  = 1
	KB = B * 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
)

func main() {
	fmt.Fprintf(os.Stdout, "%v\n", KB)
	fmt.Fprintf(os.Stdout, "%v\n", MB)
	fmt.Fprintf(os.Stdout, "%v\n", GB)
}
