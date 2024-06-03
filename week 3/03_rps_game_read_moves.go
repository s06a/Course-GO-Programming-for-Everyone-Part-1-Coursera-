/* rock paper scissors code to just read
   user's move
*/

package main

import "fmt"

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
}
