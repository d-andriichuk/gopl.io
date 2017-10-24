package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type fetchData struct {
	URL  string
	Data string
}

func main() {
	start := time.Now()
	//ch := make(chan string)
	ch := make(chan fetchData)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //run go thread
	}
	for range os.Args[1:] {
		// fmt.Println(<-ch)//recieve data from ch
		var data fetchData
		data = <-ch //recieve data from ch
		fmt.Println(data.Data)
		if ok := writeToNewFile(data); !ok {
			fmt.Fprintf(os.Stderr, "Something going wrong... Can not write data to file for URL %s\n", data.URL)
			continue
		}

	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//func fetch(url string, ch chan<- string) {
func fetch(url string, ch chan<- fetchData) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// ch <- fmt.Sprint(err)
		ch <- fetchData{URL: url, Data: fmt.Sprint(err)} //send err in ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		// ch <- fmt.Sprintf("While reading %s: %v", url, err)
		ch <- fetchData{URL: url, Data: fmt.Sprintf("While reading %s: %v", url, err)}
		return
	}
	secs := time.Since(start).Seconds()
	// ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	ch <- fetchData{URL: url, Data: fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)}
}

func writeToFile(req fetchData) bool {
	url := req.URL
	data := req.Data

	if ok := strings.HasPrefix(url, "http://"); ok {
		url = strings.TrimPrefix(url, "http://")
	}
	if ok := strings.HasPrefix(url, "https://"); ok {
		url = strings.TrimPrefix(url, "https://")
	}

	fileName := strings.Join([]string{url, ".txt"}, "")

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall - file with name %s does not exist. It will be create.\n", fileName)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetchall - save to file %s failed. Try again.\n", fileName)
			return false
		}
		defer file.Close()
		file.WriteString(data)

		return true
	}

	defer file.Close()
	file.WriteString(data)

	return true
}

func writeToNewFile(req fetchData) bool {
	url := req.URL
	data := req.Data

	if ok := strings.HasPrefix(url, "http://"); ok {
		url = strings.TrimPrefix(url, "http://")
	}
	if ok := strings.HasPrefix(url, "https://"); ok {
		url = strings.TrimPrefix(url, "https://")
	}

	fileName := strings.Join([]string{url, ".txt"}, "")

	err := os.Remove(fileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not remove the file for URL %s\n", url)
		//return false
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall - save to file %s failed. Try again.\n", fileName)
		return false
	}
	defer file.Close()
	file.WriteString(data)

	return true
}
