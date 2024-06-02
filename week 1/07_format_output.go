package main

import "fmt"

func main() {
	x := 34567.89
	fmt.Println("\nPrintln, default format i = ", x)
	fmt.Printf("Printf, f format, i = %f\n", x)
	fmt.Printf("Printf, e format, i = %e\n", x)
	fmt.Printf("Printf, 10g format, i = %10g\n", x)
	fmt.Printf("Printf, g format, i = %g\n", x)
	fmt.Printf("Printf, 10.2g format, i = %10.2g\n", x)
}
