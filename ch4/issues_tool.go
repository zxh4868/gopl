package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const Token = "xxxxxxxxxxxxxxxxxxxxxxxxxxxx"

type IssueCreate struct {
	Title string `json:"title"`
	Body  string `json:"body"` // Markdown 格式
	//Assignees []string `json:"assignees"`
	//Labels    []string `json:"labels"`
	//Milestone int `json:"milestone"`
}

type IssueReadResp struct {
	Number  int    `json:"number"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	HTMLURL string `json:"html_url"`
}

func main() {
	cmd := flag.String("cmd", "", "commond")
	owner := flag.String("owner", "", "The target owner of the issue")
	repo := flag.String("repo", "", "The target repository of the issue")
	title := flag.String("title", "", "The title of the issue")
	body := flag.String("body", "", "The body of the issue")
	number := flag.Int("number", 0, "The number of the issue")

	flag.Parse()

	switch *cmd {
	case "create", "update", "read", "lock":
	default:
		fmt.Println("Expected 'create', 'update', 'read' or 'lock' subcommands")
		os.Exit(1)
	}
	if *cmd == "create" {
		issue := IssueCreate{
			Title: *title,
			Body:  *body,
		}
		issueJson, err := json.Marshal(issue)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//fmt.Println(string(issueJson))
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", *owner, *repo)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(issueJson))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(url)
		// 设置请求头
		req.Header.Set("Authorization", "Bearer "+Token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/vnd.github+json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error executing request: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Response Status:")
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("Error creating issue: %v\n", resp.Status)
		}
		bytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(bytes))

		fmt.Println("------------------------ Issue created successfully")
	} else if *cmd == "update" {
		issue := IssueCreate{
			Title: *title,
			Body:  *body,
		}
		issueJson, err := json.Marshal(issue)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", *owner, *repo, *number)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(issueJson))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}
		req.Header.Set("Authorization", "Bearer "+Token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/vnd.github+json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error executing request: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error updating issue: %v\n", resp.Status)
			os.Exit(1)
		}
		fmt.Println("------------------------------------ Issue updated successfully")
	} else if *cmd == "read" {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d", *owner, *repo, *number)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			os.Exit(1)
		}
		req.Header.Set("Authorization", "Bearer "+Token)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/vnd.github+json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error executing request: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error reading issue: %v\n", resp.Status)
			os.Exit(1)
		}
		var issueResp IssueReadResp
		if err = json.NewDecoder(resp.Body).Decode(&issueResp); err != nil {
			fmt.Printf("Error decoding response: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Issue: %v\n", issueResp)
	} else if *cmd == "lock" {
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d/lock", *owner, *repo, *number)
		param := make(map[string]string)
		param["lock_reason"] = "off-topic"
		body, err := json.Marshal(param)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(url)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Set("Authorization", "Bearer "+Token)
		req.Header.Set("Accept", "application/vnd.github+json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusNoContent {
			fmt.Printf("Error Reading resq %v", err)
			os.Exit(1)
		}
		fmt.Printf("lock issue %d success", *number)
	}
}
