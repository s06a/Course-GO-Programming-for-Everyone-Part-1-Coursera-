package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(d []int) { // bubblesort O(n*n)
	for i := 0; i < len(d)-1; i++ {
		for j := 0; j < len(d)-i-1; j++ {
			if d[j] > d[j+1] {
				d[j], d[j+1] = d[j+1], d[j]
			}
		}
	}
}

func main() {
	startTime := time.Now()
	rand.Seed(time.Now().UnixNano())
	size := 1000
	dataSlice := make([]int, size)
	for i, _ := range dataSlice {
		dataSlice[i] = rand.Intn(size)
	}
	fmt.Println("not sorted first 20:", dataSlice[:20])
	bubbleSort(dataSlice)
	fmt.Println("first 20 after sorting:", dataSlice[:20])
	fmt.Println("it took", time.Since(startTime))
}