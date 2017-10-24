package main

import (
	"fmt"
	"gopl.io/ch1/lissajousLib"
	"log"
	"net/http"
	"strconv"
	//"strings"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	lissCycles := 0
	if c, err := strconv.Atoi(r.FormValue("cycles")); err != nil {
		lissCycles = c
	}
	lissajousLib.Lissajous(w, lissCycles)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	lissCycles := 0
	// if c, err := strconv.Atoi(r.FormValue("cycles")); err != nil {
	// 	lissCycles = c
	// }
	lissajousLib.Lissajous(w, lissCycles)
}

// func needCycles(h http.HandleFunc) http.HandleFunc {

// }
