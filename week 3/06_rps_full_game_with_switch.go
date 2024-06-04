/* a fully playable rock paper scissors game
   using switch case instead of if conditions
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
		fmt.Println("machine move is", machineMove)

		// determine result using switch
		switch move {
		case rock:
			if machineMove == rock {
				fmt.Println("-> draw")
				draws++
			} else if machineMove == paper {
				fmt.Println("-> machine wins")
				machineWins++
			} else {
				fmt.Println("-> you win")
				wins++
			}
		case paper:
			if machineMove == rock {
				fmt.Println("-> you win")
				wins++
			} else if machineMove == paper {
				fmt.Println("-> draw")
				draws++
			} else {
				fmt.Println("-> machine wins")
				machineWins++
			}
		case scissors:
			if machineMove == rock {
				fmt.Println("-> machine wins")
				machineWins++
			} else if machineMove == paper {
				fmt.Println("-> you win")
				wins++
			} else {
				fmt.Println("-> draw")
				draws++
			}
		}
	}
	fmt.Println("\nafter ", rounds, " rounds",
		"you win ", wins, ", machine wins ", machineWins,
		", with ", draws, " draws")
}
