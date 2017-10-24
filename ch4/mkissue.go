package main

import (
	"flag"
	"fmt"
	"gopl.io/ch4/github"
	"os"
)

var userName = flag.String("u", "", "user name")
var password = flag.String("p", "", "password")
var repoOwner = flag.String("ro", *userName, "repo's owner")
var repo = flag.String("r", "", "repo's name")

func main() {
	flag.Parse()
	resp, err := github.MakeIssue(*userName, *password, *repo, *userName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "mkissue: all fucked up - %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Response: %s:\n", resp)
	// fmt.Printf(format, ...)
}
