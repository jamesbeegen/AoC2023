package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Line struct {
	nums	[]num
	symbols	[]symbol
	index	int
}

type symbol struct {
	index	int
}

type num struct {
	value		string
	startIndex	int
	endIndex	int
}

func checkLine(number num, compLine Line) bool {
	adjacent := false
	for _,sym := range compLine.symbols {
		switch sym.index {
			case number.startIndex, number.startIndex-1, number.startIndex+1:
				adjacent = true
			case number.endIndex, number.endIndex-1, number.endIndex+1:
				adjacent = true
		}
	}

	return adjacent
}
func part1(challengeData []string) int {
	
	symbols := "!@#$%^&*()_+=-/"
	var vals []int
	var lines []Line

	// Get all the data into structs
	for lineIndex,l := range challengeData {
		if len(l) == 0 {
			continue
		}
		line := Line{index: lineIndex}
		for charIndex,char := range l {
			// Get the symbols
			if strings.Contains(symbols, string(char)) {
				symbol := symbol{index: charIndex}
				line.symbols = append(line.symbols, symbol)
			} else {
			// Get the numbers
				// Test if the char is a digit
				_, err := strconv.Atoi(string(char))
				if err != nil {
					continue
				}

				// Save number info to struct
				number := num {
					startIndex: charIndex, 
					endIndex: charIndex,
					value: string(char),
				}

				// Check if previous char was a digit
				if len(line.nums) == 0 {
					line.nums = append(line.nums, number)
					continue
				}
				if line.nums[len(line.nums)-1].endIndex == number.startIndex - 1 {
					line.nums[len(line.nums)-1] = num {
						value: line.nums[len(line.nums)-1].value + number.value,
						startIndex: line.nums[len(line.nums)-1].startIndex,
						endIndex: number.endIndex,
					}
				} else {
					line.nums = append(line.nums, number)
				}
			}
		}
		lines = append(lines, line)
	}
	
	for lineIndex,line := range lines {
		fmt.Println(line.nums)

		for _,number := range line.nums {
			adjacent := false
			// Check line before
			if lineIndex != 0 {
				adjacent = checkLine(number, lines[lineIndex-1])
				if adjacent {
					digit,_ := strconv.Atoi(number.value)
					vals = append(vals, digit)
					continue
				}
			}
			// Check next line
			if lineIndex != len(lines) - 1 {
				adjacent = checkLine(number, lines[lineIndex+1])
				if adjacent {
					digit,_ := strconv.Atoi(number.value)
					vals = append(vals, digit)
					continue
				}
			}
			// Check this line
			adjacent = checkLine(number, lines[lineIndex])
			if adjacent {
				digit,_ := strconv.Atoi(number.value)
				vals = append(vals, digit)
			}
		}
	}

	sum := 0
	for _,val := range vals {
		sum += val
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