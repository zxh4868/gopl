package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown 格式
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func withInOneMonth(create time.Time) bool {
	now := time.Now()
	oneMonthBefore := now.AddDate(0, -1, 0)
	return create.After(oneMonthBefore)
}

func withInOneYear(create time.Time) bool {
	now := time.Now()
	oneYearBefore := now.AddDate(-1, 0, 0)
	return create.After(oneYearBefore)
}

func test() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues found\n", len(result.Items))

	fmt.Println("With in one month:")
	for _, issue := range result.Items {
		if withInOneMonth(issue.CreatedAt) {
			fmt.Printf("#%-5d %9.9s %10.10s %.55s\n", issue.Number, issue.User.Login, issue.CreatedAt, issue.Title)
		}

	}

	fmt.Println("With in one year:")
	for _, issue := range result.Items {
		if withInOneYear(issue.CreatedAt) {
			fmt.Printf("#%-5d %9.9s %10.10s %.55s\n", issue.Number, issue.User.Login, issue.CreatedAt, issue.Title)
		}

	}
}
