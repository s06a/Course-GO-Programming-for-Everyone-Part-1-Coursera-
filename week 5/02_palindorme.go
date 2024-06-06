package main

import (
	"fmt"
	"strings"
)

func main() {
	var strInput string
	fmt.Scanf("%s", &strInput)
	strTest := strings.ToLower(strInput)
	sum := len(strTest)
	for i := 0; i < sum; i++ {
		if strTest[i] == strTest[sum-(i+1)] {
			continue
		} else {
			fmt.Println("not a palindorme")
			return
		}
	}
	fmt.Println("a palindrome!")
}