package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Map struct {
	destinationStart		int
	sourceStart 			int
	rangeLength				int
}

func getSeeds(line string) []int {
	var seeds []int

	parts := strings.Split(line, ":")
	for _, part := range strings.Split(parts[1], " ") {
		if part != "" {
			seed,_ := strconv.Atoi(part)
			seeds = append(seeds, seed)
		}
	}

	return seeds
}

func getMap(lines []string) []Map {
	maps := []Map{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		d, _ := strconv.Atoi(parts[0])
		s, _ := strconv.Atoi(parts[1])
		r, _ := strconv.Atoi(parts[2])
		m := Map{destinationStart: d, sourceStart: s, rangeLength: r}
		maps = append(maps, m)
	}

	return maps
}

func getMappedValue(source int, maps []Map) int {
	value := source
	for _, m := range maps {
		if source >= m.sourceStart && source <= (m.sourceStart + m.rangeLength) {
			diff := source - m.sourceStart
			value = m.destinationStart + diff

			fmt.Println(source, " maps to ", value)
			break
		}
	}

	return value
}

func part1(challengeData []string) int {
	var blankLines []int
	var lowestLocation int

	// Get index of blank lines
	for lineIndex, line := range challengeData {
		if len(line) == 0 {
			blankLines = append(blankLines, lineIndex)
		}
	}

	// Seed to soil map
	seedToSoilMaps := getMap(challengeData[blankLines[0]+2:blankLines[1]])

	// Soil to fertilizer maps
	soilToFertilizerMaps := getMap(challengeData[blankLines[1]+2:blankLines[2]])

	// Fertilizer to water maps
	fertilizerToWaterMaps := getMap(challengeData[blankLines[2]+2:blankLines[3]])

	// Water to light maps
	waterToLightMaps := getMap(challengeData[blankLines[3]+2:blankLines[4]])

	// Light to temperature maps
	lightToTemperatureMaps := getMap(challengeData[blankLines[4]+2:blankLines[5]])

	// Temperature to humidity maps
	temperatureToHumidityMaps := getMap(challengeData[blankLines[5]+2:blankLines[6]])

	// Humidity to location maps
	humidityToLocationMaps := getMap(challengeData[blankLines[6]+2:blankLines[7]])

	// List of seeds
	seeds := getSeeds(challengeData[0])
	for i, seed := range seeds {
		soil := getMappedValue(seed, seedToSoilMaps)
		fertilizer := getMappedValue(soil, soilToFertilizerMaps)
		water := getMappedValue(fertilizer, fertilizerToWaterMaps)
		light := getMappedValue(water, waterToLightMaps)
		temperature := getMappedValue(light, lightToTemperatureMaps)
		humidity := getMappedValue(temperature, temperatureToHumidityMaps)
		location := getMappedValue(humidity, humidityToLocationMaps)

		if i == 0 {
			lowestLocation = location
		} else {
			if location < lowestLocation {
				lowestLocation = location
			}
		}
	}

	return lowestLocation
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