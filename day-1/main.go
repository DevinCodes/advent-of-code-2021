package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	values := make([]int, 0)
	count := 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		index := len(values)
		values = append(values, num)
		if index < 1 {
			continue
		}

		prevValue := values[index-1]

		if prevValue < num {
			count++
		}
	}

	fmt.Printf("Number of increases: %d\n", count)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
