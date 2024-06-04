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
	mySlice := make([]float64, 10)
	randNum := 0.0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randNum = rand.Float64() * 100               // rand num between 0-100
		mySlice[i] = float64(int(randNum*100)) / 100 // to round the number
	}
	// find min
	fmt.Println("slice is", mySlice)
	fmt.Println("min of the slice is", minSlice(mySlice))
}
