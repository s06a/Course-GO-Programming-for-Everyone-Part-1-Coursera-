/* a code to calculate area of a circle
   using the radius read from user input */

package main

import (
	"fmt"
)

func main() {
	var radius, area float32
	const pi = 3.14159 // use math.Pi from math package instead
	fmt.Printf("input the radius: ")
	fmt.Scanf("%f", &radius)
	area = pi * radius * radius
	fmt.Printf("area is %f\n", area)
}
