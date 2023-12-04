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
	value	string
	index	int
}

type num struct {
	value		string
	startIndex	int
	endIndex	int
}

func lineData(challengeData []string) []Line {
	symbols := "!@#$%^&*()_+=-/"
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
				symbol := symbol{index: charIndex, value: string(char)}
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

	return lines
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
	var values []int
	lines := lineData(challengeData)

	for lineIndex,line := range lines {
		for _,number := range line.nums {
			adjacent := false
			// Check line before
			if lineIndex != 0 {
				adjacent = checkLine(number, lines[lineIndex-1])
				if adjacent {
					digit,_ := strconv.Atoi(number.value)
					values = append(values, digit)
					continue
				}
			}
			// Check next line
			if lineIndex != len(lines) - 1 {
				adjacent = checkLine(number, lines[lineIndex+1])
				if adjacent {
					digit,_ := strconv.Atoi(number.value)
					values = append(values, digit)
					continue
				}
			}
			// Check this line
			adjacent = checkLine(number, lines[lineIndex])
			if adjacent {
				digit,_ := strconv.Atoi(number.value)
				values = append(values, digit)
			}
		}
	}

	sum := 0
	for _,val := range values {
		sum += val
	}

	return sum
}

func checkGear(symbol symbol, compLine Line) (bool, []int) {
	adjacent := false
	var values []int
	for _,num := range compLine.nums {
		switch num.startIndex {
			case symbol.index, symbol.index-1, symbol.index+1:
				adjacent = true
				value,_ := strconv.Atoi(num.value)
				values = append(values, value)
		
		}
		switch num.endIndex {
			case symbol.index, symbol.index-1, symbol.index+1:
				adjacent = true
				value,_ := strconv.Atoi(num.value)
				if len(values) > 0 {
					if values[len(values)-1] != value {
						values = append(values, value)
					}
				} else {
					values = append(values, value)
				}
		}
	}

	return adjacent, values
}

func part2(challengeData []string) int {
	sum := 0
	lines := lineData(challengeData)

	for lineIndex,line := range lines {
		for _,symbol := range line.symbols {
			var values []int
			numAdjacent := 0

			if symbol.value != "*" {
				continue
			}

			// Check the line before
			if lineIndex != 0 {
				adjacent,vals := checkGear(symbol, lines[lineIndex-1])
				if adjacent {
					for _,val := range vals {
						values = append(values, val)
						numAdjacent++
					}
				}
			}

			// Check next line
			if lineIndex != len(lines) - 1 {
				adjacent,vals := checkGear(symbol, lines[lineIndex+1])
				if adjacent {
					for _,val := range vals {
						values = append(values, val)
						numAdjacent++
					}
				}
			}

			// Check current line
			adjacent,vals := checkGear(symbol, lines[lineIndex])
			if adjacent {
				for _,val := range vals {
					values = append(values, val)
					numAdjacent++
				}
			}
			
			if numAdjacent == 2 {
				fmt.Println("Length of values is", len(values))
				sum += values[0]*values[1]
			}
		}
	}

	return sum
}

func main() {
	// Read data from input file
	challengeData := utils.Input()

	// Print answers
	utils.Answers(strconv.Itoa(part1(challengeData)), strconv.Itoa(part2(challengeData)))
}