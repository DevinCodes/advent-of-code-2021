package main

import (
	advent_of_code "advent-of-code"
	"fmt"
)

func main() {
	values := advent_of_code.NumbersListFromInput()
	increases := advent_of_code.NumberOfIncreases(values)

	fmt.Printf("Number of increases: %d\n", increases)
}
