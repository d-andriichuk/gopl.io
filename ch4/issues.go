package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	var lessMonth []string
	var lessYear []string
	var moreYear []string
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		if time.Since(item.CreateAt).Hours() <= 744 {
			lessMonth = append(lessMonth, fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title))
		}
		if time.Since(item.CreateAt).Hours() > 744 && time.Since(item.CreateAt).Hours() <= 8928 {
			lessYear = append(lessYear, fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title))
		}
		if time.Since(item.CreateAt).Hours() > 8928 {
			moreYear = append(moreYear, fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title))
		}
	}
	fmt.Printf("%d topics:\n", result.TotalCount)
	fmt.Println("\nAdded less than a month ago:")
	for _, item := range lessMonth {
		fmt.Print(item)
	}
	fmt.Println("\nAdded less than a year ago:")
	for _, item := range lessYear {
		fmt.Print(item)
	}
	fmt.Println("\nAdded over a year ago:")
	for _, item := range moreYear {
		fmt.Print(item)
	}
}
