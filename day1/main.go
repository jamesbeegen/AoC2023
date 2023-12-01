package main

import (
	"strconv"
	"strings"
	"utils"
)

// Performs part 1 of the challenge
func part1(challengeData []string) int {
	sum := 0
	for _, line := range challengeData {
		firstInt := -1
		var lastInt int
		for _, char := range []rune(line) {
			i, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			if firstInt == -1 {
				firstInt = i
			}
			lastInt = i
		}
		sum += ((firstInt * 10) + lastInt)
	}

	return sum
}

// Performs part 2 of the challenge
func part2(challengeData []string) int {
	// New array with substituted values
	var newLines []string

	// Value mapping - preserve the first and last char
	// e.g. 'sevenine' should be replaced with 's7nine' so the 'nine' still gets replaced
	var m = map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	// Substitute values in the array
	for _, line := range challengeData {
		newLine := line
		for k, v := range m {
			newLine = strings.ReplaceAll(newLine, k, v)
		}
		newLines = append(newLines, newLine)
	}

	// Reuse part1 logic
	return part1(newLines)
}

func main() {
	// Read data from input file
	challengeData := utils.Input("input.txt")

	// Print answers
	utils.Answers(strconv.Itoa(part1(challengeData)), strconv.Itoa(part2(challengeData)))
}
