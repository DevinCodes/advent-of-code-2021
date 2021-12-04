package main

import (
	day_4 "day-4"
	"fmt"
)

func main() {
	boards := day_4.GetBoardsFromInput()
	numbers := day_4.GetNumbersFromInput()

	var lastWinningScore int
	var completedBoards []int

	for _, number := range numbers {
		for i, board := range boards {
			if contains(completedBoards, i) {
				continue
			}

			board.Mark(number)

			if board.HasBingo() {
				lastWinningScore = number * board.Score()
				completedBoards = append(completedBoards, i)
			}
		}
	}

	fmt.Printf("The score for this bingo card is %d\n", lastWinningScore)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
