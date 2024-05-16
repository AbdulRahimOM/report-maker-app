package tools

import "fmt"

func LogReport(title string,report *string) {
	fmt.Println(title,":")
	fmt.Println("===============================")
	fmt.Println(*report)
	fmt.Println("===============================")
}
