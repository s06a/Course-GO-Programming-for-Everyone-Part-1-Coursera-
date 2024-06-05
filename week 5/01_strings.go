package main

import "fmt"

func main() {
	var name string = "the loop!"
	for i := 0; i < len(name); i++ {
		if i == 1 {
			fmt.Printf("the %dst char is %c\n", i, name[i])
		} else if i == 2 {
			fmt.Printf("the %dnd char is %c\n", i, name[i])
		} else if i == 3 {
			fmt.Printf("the %drd char is %c\n", i, name[i])
		} else {
			fmt.Printf("the %dth char is %c\n", i, name[i])
		}
	}
}