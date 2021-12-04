package main

import (
	day_3 "day-3"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := day_3.LinesFromInput()

	gammaRate := day_3.GammaRate(lines)
	epsilonRate := day_3.EpsilonRate(lines)

	g, _ := strconv.ParseInt(strings.Join(gammaRate, ""), 2, 64)
	e, _ := strconv.ParseInt(strings.Join(epsilonRate, ""), 2, 64)

	fmt.Printf("Power consumption: %d\n", g*e)
}
