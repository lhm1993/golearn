//function will be used
package demo_github2

import (
	"demo_struct/demo_github1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*demo_github1.IssueSearchResult, error) {
	//将需要传递的参数使用string.Join()函数连接起来
	q := url.QueryEscape(strings.Join(terms, " "))
	//获取该网页的内容
	resp, err := http.Get(demo_github1.IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query falied :%s", resp.Status)
	}
	var result demo_github1.IssueSearchResult
	//将获得的json对象转换成IssueSearchResult类型的结构体
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	//关闭流
	resp.Body.Close()
	return &result, nil
}
