package day_2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	Direction string
	Amount int
}

func OperationListFromInput() []Operation {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ops []Operation
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		op := Operation{
			Direction: parts[0],
			Amount:    amount,
		}
		ops = append(ops, op)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ops
}
