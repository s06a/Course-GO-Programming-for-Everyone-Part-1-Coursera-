package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var mySlice []int // slice is flexible
	mySecondSlice := make([]int, 5)
	var myArray [1]int // array is more restricted
	var mySecondArray = [2]float32{5.0, 0}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, rand.Intn(100))
		mySecondSlice = append(mySecondSlice, rand.Intn(100))
	}
	myArray[0] = mySlice[0]
	mySecondArray[1] = float32(mySlice[1])
	fmt.Println("myArray is", myArray)
	fmt.Println("mySlice is", mySlice)
	fmt.Println("mySecondSlice is", mySecondSlice)
	fmt.Println("len of second slice is", len(mySecondSlice))
	fmt.Println("mySecondArray is", mySecondArray)
}