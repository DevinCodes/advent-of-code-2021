package day_1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func NumbersListFromInput() []int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var values []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
}
