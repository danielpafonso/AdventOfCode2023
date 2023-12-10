package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func parseCalibrateData(data []string) []int {
	calibrate := make([]int, 0)
	for _, line := range data {
		numbers := make([]int, 0)
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
				"Solution 1 for day 1 puzzle.\n\n" +
				"positional arguments:\n" +
				"  FILE         path for puzzle input file",
		)
		os.Exit(1)
	}
	filepath := os.Args[1]
	data := readInput(filepath)
	calibrate := parseCalibrateData(data)
	sumData(calibrate)
}
