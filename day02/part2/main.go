package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id    int
	red   int
	green int
	blue  int
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func parseGames(lines []string) []Game {
	gamesData := []Game{}
	for _, line := range lines {
		// create game with id
		split := strings.Split(line, ": ")
		id, _ := strconv.Atoi(split[0][5:])
		data := Game{id: id}
		// parse games and get maximum number of dices
		for _, game := range strings.Split(split[1], "; ") {
			for _, dice := range strings.Split(game, ", ") {
				sdice := strings.Split(dice, " ")
				numDie, _ := strconv.Atoi(sdice[0])

				if sdice[1] == "red" && data.red < numDie {
					data.red = numDie
				}
				if sdice[1] == "green" && data.green < numDie {
					data.green = numDie
				}
				if sdice[1] == "blue" && data.blue < numDie {
					data.blue = numDie
				}
			}
		}
		gamesData = append(gamesData, data)
	}
	return gamesData
}

func powerCubes(games []Game) []int {
	power := make([]int, 0)
	for _, game := range games {
		power = append(power, game.red*game.green*game.blue)
	}
	return power
}

func main() {
	// check for,required, input file
	if len(os.Args) == 1 {
		fmt.Println(
			"usage: go run solve1.go FILE\n\n" +
				"Solution 2 for day 2 puzzle.\n\n" +
				"positional arguments:\n" +
				"  FILE         path for puzzle input file",
		)
		os.Exit(1)
	}
	filepath := os.Args[1]
	data := readInput(filepath)
	games := parseGames(data)
	power := powerCubes(games)

	sum := 0
	for _, value := range power {
		sum += value
	}
	fmt.Printf("Sum of Games' power: %d\n", sum)
}
