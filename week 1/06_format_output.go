package main

import "fmt"

func main() {
	i := 55
	fmt.Println("\nPrintln, default format i = ", i)
	fmt.Printf("Printf, decimal format, i = %d\n", i)
	fmt.Printf("Printf, width 10d format, i = %10d\n", i)
	fmt.Printf("Printf, binary format, i = %b\n", i)
	fmt.Printf("Printf, octal format, i = %o\n", i)
	fmt.Printf("Printf, hexadecimal format, i = %x\n", i)
}
