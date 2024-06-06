package main

import (
	"fmt"
	"os"
)

func checkFile(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Create("06_data")
	checkFile(err)
	defer file.Close()
	var dataSlice = []int{45, 44, 84684, 21}
	for _, v := range dataSlice {
		_, err = fmt.Fprintf(file, "%d\n", v)
		checkFile(err)
	}
	fmt.Println(dataSlice)
}