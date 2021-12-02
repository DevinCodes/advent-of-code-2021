package main

import (
	day_2 "day-2"
	"fmt"
)

func main() {
	operations := day_2.OperationListFromInput()

	sub := Submarine{}

	for _, operation := range operations {
		switch operation.Direction {
		case "forward":
			sub.xPos += operation.Amount
			break
		case "down":
			sub.yPos -= operation.Amount
			break
		case "up":
			sub.yPos += operation.Amount
			break
		}
	}

	fmt.Println(sub)
	fmt.Printf("Solution: %d", sub.xPos * (-1 * sub.yPos))
}

type Submarine struct {
	xPos int
	yPos int
}