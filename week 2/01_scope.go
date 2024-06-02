package main

import "fmt"

var tempVar = "global" // camel case

func main() {
	fmt.Printf("tempVar is a %s variable\n", tempVar)
	{
		tempVar := "local"
		fmt.Printf("tempVar is a %s variable\n", tempVar)
	}
	fmt.Printf("tempVar is a %s variable\n", tempVar)
}
