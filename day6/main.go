package main

import (
	"strconv"
	"strings"
	"utils"
)


func getLineData(line string, part2 bool) []int {
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

	if part2 {
		// Append all individual data to a single data unit
		newData := ""
		for _,data := range datas {
			newData += strconv.Itoa(data)
		}

		// Clear the original datas slice
		datas = nil

		// Append the single data unit to the slice
		singleInt, _ := strconv.Atoi(newData)
		datas = append(datas, singleInt)
	}

	return datas
}

func solve(challengeData []string, part2 bool) int {
	// Get the times and distances
	times := getLineData(challengeData[0], part2)
	distances := getLineData(challengeData[1], part2)

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

func main() {
	// Read data from input file
	challengeData := utils.Input()

	// Print answers
	utils.Answers(strconv.Itoa(solve(challengeData, false)), strconv.Itoa(solve(challengeData, true)))
}