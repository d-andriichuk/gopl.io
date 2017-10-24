package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	// "net/url"
	"os"
	"strings"
)

type comics struct {
	Month      string `json:"month,omitempty"`
	Num        int    `json:"num,omitempty"`
	Link       string `json:"link,omitempty"`
	Year       string `json:"year,omitempty"`
	News       string `json:"news,omitempty"`
	SafeTitle  string `json:"safe_string,omitempty"`
	Transcript string `json:"transcript,omitempty"`
	Alt        string `json:"alt,omitempty"`
	ImageURL   string `json:"img,omitempty"`
	Title      string `json:"title,omitempty"`
	Day        string `json:"day,omitempty"`
}

const (
	defaultBaseURL = "https://xkcd.com/"

	defaultFilename = "xkcd-db.txt"
)

var cmd = flag.String("cmd", "", "Program command")

func main() {
	flag.Parse()

	var comicses []comics

	switch strings.ToUpper(*cmd) {
	case "DOWNLOAD":
		client := http.DefaultClient
		for i := 1; i < 100; i++ {
			u := fmt.Sprintf("%v%d%v", defaultBaseURL, i, "/info.0.json")
			// cURL, _ := url.Parse(u)

			req, err := http.NewRequest("GET", u, nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "xkcd: cann't create a request. Error - %v\n", err)
				os.Exit(1)
			}

			resp, err := client.Do(req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "xkcd: cann't do a request. Error - %v\n", err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "xkcd: cann't read a respons body. Error - %v\n", err)
				_, err := io.Copy(os.Stdout, resp.Body)
				if err != nil {
					fmt.Fprintf(os.Stderr, "xkcd: Fucking response. Cannt show you this shit.\n")
					os.Exit(1)
				}
			}

			var c comics
			err = json.Unmarshal(b, &c)
			if err != nil {
				fmt.Fprintf(os.Stderr, "xkcd: Cann't unmarshal this shit.\n")
				fmt.Printf("%v\n", c)
				os.Exit(1)
			}
			comicses = append(comicses, c)
		}

		data, err := json.MarshalIndent(comicses, "", " ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: Cann't marshal comicses to json. Error - %v", err)
			os.Exit(1)
		}

		var file *os.File
		// var err error
		file, err = os.Open(defaultFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: Cann't open a file. May be it does not exist.\nTrying to create\nError - %v\n", err)
			file, err = os.Create(defaultFilename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "xkcd: Cann't create a file. Error - %v", err)
				os.Exit(1)
			}
		}
		defer file.Close()

		_, err = file.Write(data)
	case "LIST":
		file, err := os.Open(defaultFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: File does not exist. Try to download some comics before.\nError - %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: Cann't get stat of the file. Error - %v\n", err)
			os.Exit(1)
		}

		data := make([]byte, stat.Size())

		_, err = file.Read(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: Cann't read a data from file. Error - %v\n", err)
			os.Exit(1)
		}

		err = json.Unmarshal(data, &comicses)
		if err != nil {
			fmt.Fprintf(os.Stderr, "xkcd: Cann't unmarshal shit from file. Error - %v\n", err)
			fmt.Fprintf(os.Stdout, "%v\n", string(data))
			os.Exit(1)
		}

		fmt.Printf("All comicses in a db:\n")
		for _, c := range comicses {
			fmt.Printf("#%d\tTitle: %s\n", c.Num, c.Title)
		}
	}
}
