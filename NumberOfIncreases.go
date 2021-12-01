package advent_of_code

func NumberOfIncreases(set []int) int {
	count := 0
	for i, num := range set {
		if i < 1 {
			continue
		}

		prevValue := set[i-1]

		if prevValue < num {
			count++
		}
	}

	return count
}