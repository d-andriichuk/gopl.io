package main

import (
	"fmt"
	//"io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if ok := strings.HasPrefix(url, "http://"); !ok {
			url = strings.Join([]string{"http://", url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("HTTP STATUS: %s\n\n", resp.Status)
		//b, err := ioutil.ReadAll(resp.Body)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}
