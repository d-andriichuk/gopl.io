package github

import (
	// "bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//SearchIssues requests GitHub.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	// fmt.Println(IssuesURL + "?q=" + q)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Request fail: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func MakeIssue(username, password, repo, repoOwner string) ([]byte, error) {
	client := &http.Client{}
	// issueURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", repoOwner, repo)
	// authURL := fmt.Sprintf("https://api.github.com/authorizations")
	issueURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", repoOwner, repo)
	var issue = MKIssue{Title: "lol",
		Body: "it's dont working",
		// Assignee:  "d-andriichuk",
		// Milestone: 0,
		// Labels:    []string{"bug", "invalid"},
	}
	jIssue, err := json.Marshal(issue)
	// fmt.Printf("%v", jIssue)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gitfuncs: all fucked up (json) - %v\n", err)
		return nil, err
	}
	data, err := json.MarshalIndent(issue, "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "JSON marshaling failed: %s\n", err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", data)
	// buf := new(bytes.Buffer)
	// _, err = buf.Write(jIssue)
	// if err != nil {
	// 	return nil, err
	// }

	req, err := http.NewRequest("POST", issueURL, strings.NewReader(string(jIssue)))
	// req, err := http.NewRequest("POST", issueURL, nil)
	// req, err := http.NewRequest("GET", issueURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gitfuncs: all fucked up - %v\n", err)
		return nil, err
	}
	// req.SetBasicAuth(username, password)
	authStr := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	req.Header.Add("Authorization", "Basic "+authStr)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	//fmt.Fprintf(os.Stdout, "%v", req.Header)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gitfuncs: all fucked up - %v\n", err)
		// return nil, err
		b, _ := ioutil.ReadAll(resp.Body)
		return b, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gitfuncs: all fucked up - %v\n", err)
		// return nil, err
	}
	return b, err
}
