/* rock paper scissors code to read user's move
   then check it with machine's move to decide
   who wins the game.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var move rune // rune stores one character

	fmt.Println("Wellcome to Rock Paper Scissors game!")
	fmt.Println("Choose your move: R, P, or S")
	fmt.Scanf("%c", &move)

	if move == 'R' { // important to use ' instead of "
		fmt.Println("your move is R")
	} else if move == 'P' {
		fmt.Println("your move is P")
	} else if move == 'S' {
		fmt.Println("your move is S")
	} else {
		fmt.Println("illegal move!")
	}

	rand.Seed(time.Now().UnixNano())
	machineMove := rand.Intn(3) // 0 for rocke, 1 for paper, 2 for scissors

	if machineMove == 0 && move == 'S' {
		fmt.Println("machine wins")
	} else if machineMove == 1 && move == 'R' {
		fmt.Println("machine wins")
	} else if machineMove == 2 && move == 'P' {
		fmt.Println("machine wins")
	} else {
		fmt.Println("draw or you win!")
	}
}
