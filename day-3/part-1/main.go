package main

import (
	day_3 "day-3"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := day_3.LinesFromInput()

	length := len(lines[0])
	numLines := len(lines)
	threshold := numLines / 2

	fmt.Printf("Length is %d\n", length)

	var numZeros int
	var gammaRate []string
	var epsilonRate []string
	for i := 0; i < length; i++ {
		numZeros = 0
		for _, line := range lines {
			if string(line[i]) == "0" {
				numZeros++
			}
		}

		if numZeros > threshold {
			gammaRate = append(gammaRate, "0")
			epsilonRate = append(epsilonRate, "1")
		} else {
			gammaRate = append(gammaRate, "1")
			epsilonRate = append(epsilonRate, "0")
		}
	}

	g, _ := strconv.ParseInt(strings.Join(gammaRate, ""), 2, 64)
	e, _ := strconv.ParseInt(strings.Join(epsilonRate, ""), 2, 64)

	fmt.Printf("Power consumption: %d\n", g*e)
}
