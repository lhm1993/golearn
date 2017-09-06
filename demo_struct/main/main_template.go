// template project main.go
package main

import (
	"demo_struct/demo_github2"
	"html/template"
	"log"
	"os"
	"time"
)

const temp1 = `{{.TotalCount}} issues:
{{range .Items}}-----------------------
Number: {{.Number}}
User: {{.User.Login}}
Title:{{.Title| printf "%.64s"}}
Age: {{.CreatedAt| daysAgo}} days
{{end}}`

var report = template.Must(template.New("issuesList").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(temp1))

func main() {
	/*
		report,err := template.New(templatename string).Funcs(template.FuncMap map[string]interface{}).Parse(template)
		//这里的检测是否是正确的模板需要手动判断，函数会返回err信息
		if err != nil {
			log.Fatal(err)
		}
		//如果使用template.Must()，Must()直接回处理，如果有错就会宕机
	*/
	//report, err := template.New("report_test").Funcs(template.FuncMap{"daysago": daysago}).Parse(temp1)

	result, err := demo_github2.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours()) / 24
}
