package main

import (
	"strconv"
	"strings"
	"utils"
)


func getLineData(line string) []int {
	var datas []int

	// Split line by colon
	parts := strings.Split(line, ":")

	// Split the numbers portion of the line
	dataParts := strings.Split(parts[1], " ")

	// Only append the items, ignore blankspace
	for _,part := range dataParts {
		data, err := strconv.Atoi(part)
		if err != nil {
			continue
		}
		datas = append(datas, data)
	}

	return datas
}

func part1(challengeData []string) int {
	// Get the times and distances
	times := getLineData(challengeData[0])
	distances := getLineData(challengeData[1])

	// Holds the number of ways to win
	var waysToWin []int

	// Loop through the races
	for i := 0; i < len(times); i++ {
		// Holds the button hold times that break the record
		var timeBounds []int

		// Current race time and record distance
		raceTime := times[i]
		recordDistance := distances[i]

		// Loop through times
		for buttonHeld := 0; buttonHeld < raceTime; buttonHeld++ {
			// Get total distance traveled at current button hold time
			distanceTraveled := (buttonHeld * (raceTime - buttonHeld))

			// Add this button hold time to list if it breaks the record
			if distanceTraveled > recordDistance {
				timeBounds = append(timeBounds, buttonHeld)
			}
		}

		// Add the number of ways to win for this race to the slice
		waysToWin = append(waysToWin, (timeBounds[len(timeBounds)-1] - timeBounds[0]) + 1)
	}

	// Calculate the total number of ways to win from slice
	total := 1
	for _, way := range waysToWin {
		total *= way
	}

	return total
}

func part2(challengeData []string) int {
	return 0
}


func main() {
	// Read data from input file
	challengeData := utils.Input()

	// Print answers
	utils.Answers(strconv.Itoa(part1(challengeData)), strconv.Itoa(part2(challengeData)))
}