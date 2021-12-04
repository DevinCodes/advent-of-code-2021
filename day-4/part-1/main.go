package main

import (
	day_4 "day-4"
	"fmt"
)

func main() {
	boards := day_4.GetBoardsFromInput()
	numbers := day_4.GetNumbersFromInput()

	game:
	for _, number := range numbers {
		for i, board := range boards {
			board.Mark(number)

			if board.HasBingo() {
				fmt.Printf("Bingo for board #%d at number %d!\n", i+1, number)
				fmt.Printf("The score for this bingo card is %d\n", board.Score() * number)
				break game
			}
		}
	}
}
