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
	games_data := []Game{}
	for _, line := range lines {
		// create game with id
		split := strings.Split(line, ": ")
		id, _ := strconv.Atoi(split[0][5:])
		data := Game{id: id}
		// parse games and get maximum number of dices
		for _, game := range strings.Split(split[1], "; ") {
			for _, dice := range strings.Split(game, ", ") {
				sdice := strings.Split(dice, " ")
				num_die, _ := strconv.Atoi(sdice[0])

				if sdice[1] == "red" && data.red < num_die {
					data.red = num_die
				}
				if sdice[1] == "green" && data.green < num_die {
					data.green = num_die
				}
				if sdice[1] == "blue" && data.blue < num_die {
					data.blue = num_die
				}
			}
		}
		games_data = append(games_data, data)
	}
	return games_data
}

func checkValidGames(games []Game) []int {
	valid_games := make([]int, 0)
	max_red := 12
	max_green := 13
	max_blue := 14

	for _, game := range games {
		fmt.Printf("%+v\n", game)
		if game.red <= max_red && game.green <= max_green && game.blue <= max_blue {
			valid_games = append(valid_games, game.id)
		}
	}

	return valid_games
}

func main() {
	// check for,required, input file
	if len(os.Args) == 1 {
		fmt.Println(
			"usage: go run solve1.go FILE\n\n" +
				"Solution 1 for day 2 puzzle.\n\n" +
				"positional arguments:\n" +
				"  FILE         path for puzzle input file",
		)
		os.Exit(1)
	}
	filepath := os.Args[1]
	data := readInput(filepath)
	games := parseGames(data)
	valid := checkValidGames(games)

	sum := 0
	for _, value := range valid {
		sum += value
	}
	fmt.Printf("Sum of valid Games ID's: %d\n", sum)
}
