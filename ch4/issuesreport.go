package main

import (
	"gopl/ch4/github"
	th "html/template"
	"log"
	"os"
	tt "text/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
	{{range .Items}}
--------------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Time: {{.CreatedAt | daysAgo}} days
	{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	report := tt.Must(tt.New("report").Funcs(tt.FuncMap{"daysAgo": daysAgo}).Parse(templ))
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

	issueList := th.Must(th.New("issue_list").Funcs(th.FuncMap{"daysAgo": daysAgo}).Parse(`
		<h1>{{.TotalCount}} issues</h1>
		<table>
		<tr>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
			<td>{{.Number}}</td>
			<td>{{.State}}</td>
			<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
			<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
		</table>
	`))
	fileName := "issues_report.html"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	issueList.Execute(file, result)
	defer file.Close()
}
