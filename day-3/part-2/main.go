package main

import (
	day_3 "day-3"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	lines := day_3.LinesFromInput()

	oxGenNum := make([]string, len(lines))
	co2RateNum := make([]string, len(lines))
	copy(oxGenNum, lines)
	copy(co2RateNum, lines)

	length := len(oxGenNum[0])
	var zeroIndices, oneIndices, indicesToRemove []int
	for i := 0; i < length; i++ {
		if len(oxGenNum) == 1 {
			break
		}

		zeroIndices = []int{}
		oneIndices = []int{}
		indicesToRemove = []int{}
		for i2, s := range oxGenNum {
			if s[i] == '0' {
				zeroIndices = append(zeroIndices, i2)
			} else {
				oneIndices = append(oneIndices, i2)
			}
		}

		if len(zeroIndices) > len(oneIndices) {
			indicesToRemove = oneIndices
		} else {
			indicesToRemove = zeroIndices
		}

		sort.Sort(sort.Reverse(sort.IntSlice(indicesToRemove)))
		for _, i := range indicesToRemove {
			oxGenNum = remove(oxGenNum, i)
		}
	}

	length = len(co2RateNum[0])
	for i := 0; i < length; i++ {
		if len(co2RateNum) == 1 {
			break
		}

		zeroIndices = []int{}
		oneIndices = []int{}
		indicesToRemove = []int{}
		for i2, s := range co2RateNum {
			if s[i] == '1' {
				zeroIndices = append(zeroIndices, i2)
			} else {
				oneIndices = append(oneIndices, i2)
			}
		}

		if len(zeroIndices) < len(oneIndices) {
			indicesToRemove = oneIndices
		} else {
			indicesToRemove = zeroIndices
		}

		sort.Sort(sort.Reverse(sort.IntSlice(indicesToRemove)))
		for _, i := range indicesToRemove {
			co2RateNum = remove(co2RateNum, i)
		}
	}

	ox, _ := strconv.ParseInt(oxGenNum[0], 2, 64)
	co2, _ := strconv.ParseInt(co2RateNum[0], 2, 64)

	fmt.Printf("Power consumption: %d\n", ox*co2)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
