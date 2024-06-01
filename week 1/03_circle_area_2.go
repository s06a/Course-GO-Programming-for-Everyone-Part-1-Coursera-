/* a code to calculate area of a circle
   using the radius read from user input */

package main

import (
	"fmt"
	"math"
)

func main() {
	var radius, area float32
	fmt.Printf("input the radius: ")
	fmt.Scanf("%f", &radius)
	area = math.Pi * radius * radius
	fmt.Printf("area is %f\n", area)
}
