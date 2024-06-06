package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("05_data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var item int
	var dataSlice []int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &item)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		dataSlice = append(dataSlice, item)
	}
	fmt.Println(dataSlice)
}