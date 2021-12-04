package day_3

func EpsilonRate(lines []string) []string {
	length := len(lines[0])
	numLines := len(lines)
	threshold := numLines / 2

	var numZeros int
	var gammaRate []string
	for i := 0; i < length; i++ {
		numZeros = 0
		for _, line := range lines {
			if string(line[i]) == "0" {
				numZeros++
			}
		}

		if numZeros < threshold {
			gammaRate = append(gammaRate, "0")
		} else {
			gammaRate = append(gammaRate, "1")
		}
	}

	return gammaRate
}
