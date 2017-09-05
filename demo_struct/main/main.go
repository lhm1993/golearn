// demo_struct project main.go
package main

import (
	"demo_struct/demo_github2"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args)
	result, err := demo_github2.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Issues\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("%#-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
