/* a code to find min in a slice, which
   runs in O(n) -- linear time.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func minSlice(slice []float64) float64 {
	var min float64 = 1000.0
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	size := 0
	fmt.Println("input a large number as size of slice")
	fmt.Scanf("%d", &size)
	mySlice := make([]float64, size)
	randNum := 0.0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		randNum = rand.Float64() * 100                  // rand num between 0-100
		mySlice[i] = float64(int(randNum*100000)) / 100 // to round the number
	}
	// find min
	startTime := time.Now()
	fmt.Println("min of the slice is", minSlice(mySlice))
	duration := time.Since(startTime)
	fmt.Println("running time is:", duration)
}
