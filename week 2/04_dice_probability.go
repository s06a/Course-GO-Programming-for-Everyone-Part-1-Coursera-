/* Simple kind of monte carlo to calculate
   probability of getting a pair value from
   two dice.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var pair, howMany, steps int = 0, 0, 1000000
	var vPair int
	fmt.Println("reading pair of dice value")
	fmt.Scanf("%d", &vPair)
	for i := 0; i < steps; i++ {
		pair = rand.Intn(6) + rand.Intn(6) + 2
		if pair == vPair {
			howMany++
		}
	}
	fmt.Printf("probability of rolling a %d is %2.2f%%\n", vPair, float32(howMany)/float32(steps)*100)
}
