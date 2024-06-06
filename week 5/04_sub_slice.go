package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var slice = []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		slice = append(slice, rand.Intn(10))
	}
	fmt.Println("slice is:", slice)
	fmt.Println("till 3nd:", slice[:2])
	fmt.Println("till 2nd till 7th:", slice[2:7])
	fmt.Println("last:", slice[len(slice)-1:])
}