// demo_struct project main.go
package main

import (
	"demo_struct/demo_github1"
	"demo_struct/demo_github2"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

//var sorted_items map[int]*demo_github1.Issue  //这是一个空的map，不能存储数据，需要赋初值，或者使用make函数

func main() {
	fmt.Println(os.Args)
	result, err := demo_github2.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Issues\n", result.TotalCount)
	var dates []int
	sorted_items := make(map[int]*demo_github1.Issue) //新建一个map使用的两种方式
	var i int = 1
	for _, item := range result.Items {
		i++
		date := int(time.Since(item.CreatedAt).Hours()) / 24
		//fmt.Printf("%d\n", int(time.Since(item.CreatedAt).Hours())/24)
		sorted_items[date] = item
		dates = append(dates, date)

		//fmt.Printf("%#-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("i=", i)
	sort.Ints(dates)
	fmt.Println(dates)
	for _, date := range dates {
		//fmt.Printf("%-5d\n", date)
		fmt.Printf("%#-5d %9.9s %.55s %-5d\n", sorted_items[date].Number, sorted_items[date].User.Login, sorted_items[date].Title, date)
	}

}
