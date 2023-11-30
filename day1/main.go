package main

import (
	"fmt"
	"os"
	"strings"
)

// Reads the file into a string array
func readFile(fileName string) []string {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(fileContent), "\n")
}

// Performs part 1 of the challenge
func part1(challengeData []string) string {
	return ""
}

// Performs part 2 of the challenge
func part2(challengeData []string) string {
	return ""
}

func main() {
	// Read data from input file
	challengeData := readFile("input.txt")

	// Get the answers
	answer1 := part1(challengeData)
	answer2 := part2(challengeData)

	// Print answers
	fmt.Println("Part 1 Answer:", answer1)
	fmt.Println("Part 2 Answer:", answer2)
}