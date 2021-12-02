package main

import (
	day_1 "day-1"
	"fmt"
)

func main() {
	numbers := day_1.NumbersListFromInput()
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

	increases := day_1.NumberOfIncreases(sums)
	fmt.Printf("Number of increases: %d\n", increases)

}
