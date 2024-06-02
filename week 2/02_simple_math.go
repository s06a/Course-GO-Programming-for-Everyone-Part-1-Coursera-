package main

import "fmt"

func main() {
	var i, j, k int32 = 1, 2, 3
	fmt.Printf("i * j + k * 99%%k = %f\n", float32(i*j+k*99%k))
	fmt.Printf("-(i + j) = %d\n", -(i + j))
}
