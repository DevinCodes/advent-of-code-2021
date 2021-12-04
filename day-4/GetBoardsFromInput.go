package day_4

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	rows []boardRow
}

type field struct {
	value int
	marked bool
}

type boardRow []field

func (f *field) mark(num int) {
	if f.value == num {
		f.marked = true
	}
}

func (b *Board) Score() int {
	var result int

	for _, row := range b.rows {
		for _, num := range row {
			if !num.marked {
				result += num.value
			}
		}
	}

	return result
}

func (b *Board) Mark(val int) {
	for _, row := range b.rows {
		for i2, num := range row {
			num.mark(val)
			row[i2] = num
		}
	}
}

func (b *Board) HasBingo() bool {
	var lineBingo bool
	for _, row := range b.rows {
		lineBingo = true
		for _, num := range row {
			lineBingo = lineBingo && num.marked
		}

		if lineBingo {
			return true
		}
	}

	// column bingo

	var columnBingo bool
	for i := 0; i < len(b.rows[0]); i++ {
		columnBingo = true
		for _, row := range b.rows {
			columnBingo = columnBingo && row[i].marked
		}

		if columnBingo {
			return true
		}
	}

	return false
}


func GetBoardsFromInput() []Board {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var boards []Board
	var rows []boardRow
	var rowValues boardRow

	firstLine := true
	for scanner.Scan() {
		if firstLine {
			firstLine = false
			continue
		}
		if scanner.Text() == "" {
			if len(rows) > 0 {
				boards = append(boards, Board{rows: rows})
				rows = []boardRow{}
			}
			continue
		}

		rowValues = []field{}

		nums := strings.Split(scanner.Text(), " ")
		var val int
		for _, num := range nums {
			val, err = strconv.Atoi(num)
			if err != nil {
				continue
			}

			rowValues = append(rowValues, field{value: val})
		}
		rows = append(rows, rowValues)
	}

	if len(rows) > 0 {
		boards = append(boards, Board{rows: rows})
	}

	return boards
}