package main

import "fmt"

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	} else if n <= 20 { // protection for int64 to prevent negative result
		return (n * factorial(n-1))
	} else {
		panic("no more than 20!")
	}
}

func main() {
	var n int64
	fmt.Println("input number")
	fmt.Scanf("%d", &n)
	fmt.Println("result is", factorial(n))
}