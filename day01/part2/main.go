package main

import (
	"fmt"
	"os"
	"strings"
)

var numberMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

var deduplicate = map[string]string{
	"twone":     "twoone",
	"eightwo":   "eighttwo",
	"eighthree": "eightthree",
	"oneight":   "oneeight",
	"threeight": "threeeight",
	"fiveight":  "fiveeight",
	"nineight":  "nineeight",
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func translateNumber(line string) string {
	// First we must duplicate values since the spelling be my overlaped
	// This is something that isn't present in the examples
	// 85dntjeightwom -> should result int the numbers: [8 5 8 2]
	for o, n := range deduplicate {
		line = strings.ReplaceAll(line, o, n)
	}
	// replace number string with number
	for s, d := range numberMap {
		line = strings.ReplaceAll(line, s, d)
	}
	return line
}

func parseCalibrateData(data []string) []int {
	calibrate := make([]int, 0)
	for _, line := range data {
		numbers := make([]int, 0)
		line := translateNumber(line)
		for _, ascii := range line {
			if ascii >= 48 && ascii <= 57 {
				numbers = append(numbers, int(ascii)-48)
			}
		}
		value := numbers[0]*10 + numbers[len(numbers)-1]
		calibrate = append(calibrate, value)
	}
	return calibrate
}

func sumData(data []int) {
	sum := 0
	for _, i := range data {
		sum += i
	}
	fmt.Printf("Sum of Calibration values: %d\n", sum)
}

func main() {
	// check for,required, input file
	if len(os.Args) == 1 {
		fmt.Println(
			"usage: go run solve1.go FILE\n\n" +
				"Solution 2 for day 1 puzzle.\n\n" +
				"positional arguments:\n" +
				"  FILE         path for puzzle input file",
		)
		os.Exit(1)
	}
	filepath := os.Args[1]
	data := readInput(filepath)
	calibrate := parseCalibrateData(data)
	// fmt.Println(calibrate)
	sumData(calibrate)
}
