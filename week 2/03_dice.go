package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceRoll := rand.Intn(6) + 1
	fmt.Printf("You rolled: %d\n", diceRoll)
}
