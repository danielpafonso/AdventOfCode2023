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
}

var (
	numbers = map[int][]Number{}
	symbols = make([]Symbol, 0)
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
				symbols = append(symbols, Symbol{row: row, column: col})
			}
		}
	}
}

func main() {
	// check for,required, input file
	if len(os.Args) == 1 {
		fmt.Println(
			"usage: go run solve1.go FILE\n\n" +
				"Solution 1 for day x puzzle.\n\n" +
				"positional arguments:\n" +
				"  FILE         path for puzzle input file",
		)
		os.Exit(1)
	}
	filepath := os.Args[1]
	data := readInput(filepath)
	parseData(data)
	fmt.Printf("%+v\n", symbols)
	for k, v := range numbers {
		fmt.Println(k, v)
	}
}
