package main

import "fmt"

func convertToKm(miles int, yards int) float64 {
	return (1.609*float64(miles) + float64(yards)/1760.0)
}

func main() {
	var miles, yards int
	for i := 0; i < 10; i++ {
		fmt.Println("Input miles and yards as two integers")
		fmt.Println("if miles is negative, program exits")
		fmt.Scanf("%d%d\n", &miles, &yards)
		if miles < 0 {
			return
		}
		fmt.Printf("Answer is %f kilometers.\n\n", convertToKm(miles, yards))
	}
}