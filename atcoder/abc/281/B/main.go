package main

import (
	"fmt"
	"regexp"
)

func main() {
	var s string
	fmt.Scan(&s)

	pattern := "^[A-Z]([1-9][0-9]{5})[A-Z]$"
	match, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err)
	}
	if match {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
