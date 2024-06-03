package main

import "fmt"

func main() {
	i := 50
	for { // no conditions, so it's always true
		fmt.Println("i is ", i)
		i--
		if i < 45 {
			fmt.Println("terminated!")
			break
		}
	}
}
