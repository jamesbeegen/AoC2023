package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type game struct {
	id			int
	draws 		[]string
	possible 	bool
}

func possible(game game) bool {
	return false
}

func games(challengeData []string) []game {
	var games []game
	for _,line := range challengeData {
		var game game
		var err error

		if len(line) == 0 {
			continue
		}
		
		// Split by semicolon and get the Id
		parts := strings.Split(line, ":")
		subParts := strings.Split(parts[0], " ")
		id, err := strconv.Atoi(subParts[1])
		if err != nil {
			panic(err)
		}

		// Get the draws
		draws := strings.Split(parts[1], ";")
		for i,_ := range draws {
			draws[i] = strings.ReplaceAll(draws[i], " ", "")
		}
		// Instantiate game struct
		game.id = id
		game.draws = draws
		game.possible = false

		// Add the game to list
		games = append(games, game)
	}

	return games
}

func part1(challengeData []string) int {
	var games = games(challengeData)
	var possibleIds []int
	// Max values 
	var maxes = map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	// Loop through games
	for _,game := range games {
		possible := true
		for _,draw := range game.draws {
			vals := strings.Split(draw, ",")
			for _,val := range vals {
				if strings.Contains(val, "blue") {
					for i,c := range []rune(val) {
						if c == 'b' {
							numDrawn, _ := strconv.Atoi(val[:i])
							fmt.Println(numDrawn)
							if numDrawn > maxes["blue"] {
								possible = false
							}
						}
					}
				} else if strings.Contains(val, "red") {
					for i,c := range []rune(val) {
						if c == 'r' {
							numDrawn, _ := strconv.Atoi(val[:i])
							if numDrawn > maxes["red"] {
								possible = false
							}
						}
					}
				} else {
					for i,c := range []rune(val) {
						if c == 'g' {
							numDrawn, _ := strconv.Atoi(val[:i])
							if numDrawn > maxes["green"] {
								possible = false
							}
						}
					}
				}
			}
		}

		// If the game is possible, add its id to a list
		if possible {
			possibleIds = append(possibleIds, game.id)
		}
	}

	sum := 0
	for _,id := range possibleIds {
		sum += id
	}
	
	return sum
}

func part2(challengeData []string) int {
	return 1
}

func main() {
	// Read data from input file
	challengeData := utils.Input()

	// Print answers
	utils.Answers(strconv.Itoa(part1(challengeData)), strconv.Itoa(part2(challengeData)))
}