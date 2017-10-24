package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

var words = make(map[string]map[string]int)

func main() {
	for _, fName := range os.Args[1:] {
		b, err := ioutil.ReadFile(fName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq: Cann't read the file %s. Error: %v", fName, err)
			continue
		}
		if err = calcWords(fName, b); err != nil {
			fmt.Fprintf(os.Stderr, "wordfreq: Cann't scan a words from this file: %s. Error: %v", fName, err)
			continue
		}
	}
	// fmt.Println(words)
	printStat()
}

func calcWords(fName string, b []byte) error {
	reader := bytes.NewReader(b)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		addWord(fName, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func addWord(fName, word string) {
	t := words[fName]
	if t == nil {
		t = make(map[string]int)
		words[fName] = t
	}
	t[word]++
}

func printStat() {
	for fName, text := range words {
		fmt.Println(fName)
		fmt.Printf("word\tnumber\n")
		for word, num := range text {
			fmt.Printf("%s\t%d\n", word, num)
		}
	}
}
