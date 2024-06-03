package main

import "fmt"

func main() {
	i := 3 // to test the scope
	if i < 5 {
		fmt.Printf("%d is lower than 5\n", i)
	}
	fmt.Println("second if conditions")
	if i := 6; i < 5 {
		fmt.Printf("%d is lower than 5\n", i)
	} else {
		fmt.Printf("%d is greater than 5\n", i)
	}
}
