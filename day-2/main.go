package main

import (
	advent_of_code "advent-of-code"
	"fmt"
)

func main() {
	numbers := advent_of_code.NumbersListFromInput()
	var sums []int
	var currSum int
	maxIndex := len(numbers) - 1

	for i, number := range numbers {
		if i > maxIndex-2 {
			break
		}

		currSum += number
		currSum += numbers[i+1]
		currSum += numbers[i+2]

		sums = append(sums, currSum)
		currSum = 0
	}

	increases := advent_of_code.NumberOfIncreases(sums)
	fmt.Printf("Number of increases: %d\n", increases)

}
