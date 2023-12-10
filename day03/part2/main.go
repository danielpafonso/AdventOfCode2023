package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	id    int
	value int
	row   int
	start int
	end   int
}

type Symbol struct {
	row    int
	column int
	symbol string
}

type Position struct {
	row    int
	column int
}

var (
	numbers   = map[int][]Number{}
	symbols   = make([]Symbol, 0)
	positions = []Position{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
)

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func parseData(data []string) {
	numID := 0
	for row, line := range data {
		buffer := ""
		start := -1
		// add final dot to line
		line += "."
		for col, r := range line {
			char := string(r)
			// check if number
			if r >= 48 && r <= 57 {
				buffer += char
				if start == -1 {
					start = col
				}
				// process next caracter
				continue
			}
			// check if buffer have number
			if buffer != "" {
				value, _ := strconv.Atoi(buffer)
				number := Number{
					id:    numID,
					value: value,
					row:   row,
					start: start,
					end:   start + len(buffer) - 1,
				}
				numID += 1
				// add to map
				_, ok := numbers[row]
				if ok {
					numbers[row] = append(numbers[row], number)
				} else {
					numbers[row] = []Number{number}
				}
				// reset buffer and start
				buffer = ""
				start = -1
			}
			if char == "." {
				// skip if dot
				continue
			} else {
				// symbol
				symbols = append(symbols, Symbol{row: row, column: col, symbol: char})
			}
		}
	}
}

func getGearRatio() []int {
	gearRatios := make([]int, 0)

	for _, sym := range symbols {
		if sym.symbol != "*" {
			continue
		}
		distinctNumbers := map[int]int{}
		// get around positions
		for _, pos := range positions {
			checkRow := sym.row + pos.row
			checkColumn := sym.column + pos.column
			// check if there is a number on that row
			if nums, ok := numbers[checkRow]; ok {
				// check if there is a number with that column
				for _, num := range nums {
					if checkColumn >= num.start && checkColumn <= num.end {
						// insert into to distinct columns
						if _, ok := distinctNumbers[num.id]; !ok {
							distinctNumbers[num.id] = num.value
						}
					}
				}
			}
		}
		if len(distinctNumbers) == 2 {
			ratio := 1
			for _, part := range distinctNumbers {
				ratio *= part
			}
			gearRatios = append(gearRatios, ratio)
		}
	}
	return gearRatios
}

func main() {
	// check for,required, input file
	if len(os.Args) == 1 {
		fmt.Println(
			"usage: go run solve1.go FILE\n\n" +
				"Solution 1 for day 3 puzzle.\n\n" +
				"positional arguments:\n" +
				"  FILE         path for puzzle input file",
		)
		os.Exit(1)
	}
	filepath := os.Args[1]
	data := readInput(filepath)
	parseData(data)

	parts := getGearRatio()
	sum := 0
	for _, v := range parts {
		sum += v
	}
	fmt.Printf("Sum of Gear Ratios: %d\n", sum)
}
