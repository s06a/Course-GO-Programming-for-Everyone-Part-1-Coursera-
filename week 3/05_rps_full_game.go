/* a fully playable rock paper scissors game,
   which uses almost any topic we have learned
   till now such as loops, conditions, different
   variables and ...
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var move, machineMove int
	const (
		rock     = 0
		paper    = 1
		scissors = 2
	)
	const (
		cRock     = 'R'
		cPaper    = 'P'
		cScissors = 'S'
	)
	var cMove rune
	var draws, wins, machineWins int
	var rounds int

	fmt.Println("How many rounds do you want to play?")
	fmt.Scanf("%d", &rounds)

	for i := 0; i < rounds; i++ {
		// user plays
		fmt.Println("\nRound ", i+1, ": Choose either R, P, or S")
		fmt.Scanf("%c\n", &cMove)

		if cMove == cRock {
			move = rock
		} else if cMove == cPaper {
			move = paper
		} else if cMove == cScissors {
			move = scissors
		} else {
			fmt.Println("-> illegal move")
			i--
			continue
		}

		// machine plays
		rand.Seed(time.Now().UnixNano())
		machineMove = rand.Intn(3)
		if machineMove == rock {
			cMove = cRock
		} else if machineMove == paper {
			cMove = cPaper
		} else if machineMove == scissors {
			cMove = cScissors
		}
		fmt.Printf("machine plays %c\n", cMove)

		// determine result
		if move == machineMove {
			fmt.Println("-> draw")
			draws++
		} else if (move == paper) && (machineMove == scissors) {
			fmt.Println("-> machine wins")
			machineWins++
			continue
		} else if (move == rock) && (machineMove == paper) {
			fmt.Println("-> machine wins")
			machineWins++
			continue
		} else if (move == scissors) && (machineMove == rock) {
			fmt.Println("-> machine wins")
			machineWins++
			continue
		} else {
			fmt.Println("-> you win")
			wins++
		}
	}
	fmt.Println("\nafter ", rounds, " rounds",
		"you win ", wins, ", machine wins ", machineWins,
		", with ", draws, " draws")
}