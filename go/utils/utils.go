package utils

import (
	"fmt"
	"os"
	"strings"
)

// Reads the file into a string array
func Input(fileName string) []string {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(fileContent), "\n")
}

// Prints the answers to the screen
func Answers(p1 string, p2 string) {
	fmt.Println("Part 1 Answer:", p1)
	fmt.Println("Part 2 Answer:", p2)
}