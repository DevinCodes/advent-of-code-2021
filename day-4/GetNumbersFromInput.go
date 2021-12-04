package day_4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetNumbersFromInput() []int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []int
	for scanner.Scan() {
		// We only want the first line
		if scanner.Text() == "" {
			break
		}

		nums := strings.Split(scanner.Text(), ",")

		var val int
		for _, num := range nums {
			val, _ = strconv.Atoi(num)
			values = append(values, val)
		}
	}

	return values
}
