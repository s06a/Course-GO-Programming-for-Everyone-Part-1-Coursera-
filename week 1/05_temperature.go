package main

import "fmt"

func main() {
	var from, to, by int32 = 0, 250, 10
	var fahrenheit, centigrade float32
	fahrenheit = float32(from)
	for int32(fahrenheit) <= to {
		centigrade = 5.0 / 9.0 * float32(fahrenheit-32)
		fmt.Printf("%g f \t %g c\n", fahrenheit, centigrade)
		fahrenheit += float32(by)
	}
}
