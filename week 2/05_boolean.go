package main

import "fmt"

func main() {
	var p, q, r = true, false, true
	fmt.Println("p = q", p == q)
	fmt.Println("p = q = r", p == q && p == r)
	fmt.Println("p = !r", p == !r)
	fmt.Println("p = q or p = r", p == q || p == r)
}
