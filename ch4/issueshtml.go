package main

import (
	"html/template"
	"gopl.io/ch4/github"
	"log"
	"os"
)

var issuesList = template.Must(template.New("issueslist").Parse(`
	<h1>{{.TotalCount}} topics</h1>
	<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issuesList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}