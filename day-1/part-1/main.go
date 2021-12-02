package main

import (
	day_1 "day-1"
	"fmt"
)

func main() {
	values := day_1.NumbersListFromInput()
	increases := day_1.NumberOfIncreases(values)

	fmt.Printf("Number of increases: %d\n", increases)
}
