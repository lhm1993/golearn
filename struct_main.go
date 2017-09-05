// demo_struct project main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color,omitempty"`
	Actor []string
}

func main() {
	var movies = []Movie{
		{
			Title: "title1",
			Year:  1942,
			Color: false,
			Actor: []string{"1", "2"},
		}, //结构体变量初始化赋值的时候需要加上 , ，结构体数组也是如此，要加上 , ；
		{
			Title: "title2",
			Year:  1888,
			Color: true,
			Actor: []string{"1", "2", "3"},
		},
	}
	date, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", date)
	date, err = json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("json marshalindent filed:%s", err)
	}
	fmt.Printf("%s\n", date)
	var titles []struct{ Titles string }
	if err := json.Unmarshal(date, &titles); err != nil {
		log.Fatalf("JSON unmarshal failed :%", err)
	}
	fmt.Printf("%#v\n", titles)
	var test []string = []string{"1", "2"}
}
