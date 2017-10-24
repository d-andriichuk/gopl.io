package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    //unicode symbols number
	var utflen [utf8.UTFMax + 1]int //unicode lenght
	invalid := 0
	numbers := 0
	letters := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if r == rune('q') {
			break
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if ok := unicode.IsLetter(r); ok {
			letters++
		}
		if ok := unicode.IsNumber(r); ok {
			numbers++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\nLetters: %d; Numbers: %d\n", letters, numbers
		)
	if invalid > 0 {
		fmt.Printf("\nInvalid UTF-8 symbols - %d\n", invalid)
	}
}
