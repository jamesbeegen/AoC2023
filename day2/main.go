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
		for i := range draws {
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

func solve(challengeData []string, part2 bool) int {
	// Get the games into structs
	var games = games(challengeData)

	// Array holding ids of possible games
	var possibleIds []int

	// Max values 
	var maxes = map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}

	// Holds the powers for part2
	var powers []int

	// Loop through games
	for _,game := range games {
		// The max number that were drawn
		redMax := 0
		blueMax := 0
		greenMax := 0

		// Is the game possible
		possible := true

		// Loop through all of the draws
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
							if part2 {
								if numDrawn > blueMax {
									blueMax = numDrawn
								}
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
							if part2 {
								if numDrawn > redMax {
									redMax = numDrawn
								}
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
							if part2 {
								if numDrawn > greenMax {
									greenMax = numDrawn
								}
							}
						}
					}
				}
			}
		}
		if part2 {
			powers = append(powers, blueMax*redMax*greenMax)
		}

		// If the game is possible, add its id to a list
		if possible {
			possibleIds = append(possibleIds, game.id)
		}
	}

	sum := 0
	if part2 {
		for _,power := range powers {
			sum += power
		}
	} else {
		for _,id := range possibleIds {
			sum += id
		}
	}
	
	return sum
}

func main() {
	// Read data from input file
	challengeData := utils.Input()

	// Print answers
	utils.Answers(strconv.Itoa(solve(challengeData, false)), strconv.Itoa(solve(challengeData, true)))
}