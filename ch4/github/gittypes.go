package github

import (
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreatedAt time.Time `json:"created_at"`
	Body     string
}

type MKIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	// Assignee  string 	`json:"assignee"`
	// Milestone int	 	`json:"milestone"`
	// Labels    []string	`json:"labels"`
	// Assignees []string	`json:"assignees"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
