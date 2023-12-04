package main

import (
	"slices"
	"strconv"
	"strings"
	"utils"
)

type card struct {
	id				int
	winningNums		[]int
	scratchedNums	[]int
	numMatches		int
	instances		int
}

func createCard(line string) card {
	var card card

	// Split line by colon
	lineParts := strings.Split(line, ":")

	// Get card id
	subParts := strings.Split(lineParts[0], " ")
	id,_ := strconv.Atoi(subParts[len(subParts)-1])
	card.id = id

	// Split the numbers portion of the line
	numParts := strings.Split(lineParts[1], "|")

	// Get winning numbers
	for _,val := range strings.Split(numParts[0], " ") {
		int_val,err := strconv.Atoi(strings.Trim(val, " "))
		if err != nil {
			continue
		}
		card.winningNums = append(card.winningNums, int_val)
	}

	// Get scratched numbers
	for _,val := range strings.Split(numParts[1], " ") {
		int_val,err := strconv.Atoi(strings.Trim(val, " "))
		if err != nil {
			continue
		}
		card.scratchedNums = append(card.scratchedNums, int_val)
	}
	
	return card
}

func (card *card) matches() int {
	numMatch := 0
	// Check for matches
	for _,num := range card.scratchedNums {
		if slices.Contains(card.winningNums, num) {
			numMatch++
		}
	}

	return numMatch
}

func (card *card) value() int {
	value := 0

	// Get the number of matches
	card.numMatches = card.matches()

	// Get the value of the card
	if card.numMatches < 2 {
		value = 1 * card.numMatches
	} else {
		value = 1 
		for x := 1; x < card.numMatches; x++ {
			value *= 2
		}
	}

	return value
}

func part1(challengeData []string) int {
	sum := 0
	for _,line := range challengeData {
		if len(line) == 0 {
			continue
		}
		card := createCard(line)
		sum += card.value()
	}
	
	return sum
}

func copies(cards []card, cardsMap map[int]card, ) []card {
	var newCards []card

	for _,thisCard := range cards {
		for i := 1; i <= thisCard.numMatches; i++ {
			new_card := cardsMap[thisCard.id + i]
			newCards = append(newCards, new_card)
		}
	}
	
	return newCards
}

func part2(challengeData []string) int {
	var cards []card
	cardsMap :=  make(map[int]card)

	// Get all of the original cards
	for _,line := range challengeData {
		if len(line) == 0 {
			continue
		}

		// Create a card
		card := createCard(line)
		
		// Get number of matches
		card.numMatches = card.matches()

		// Add card to list and map
		cardsMap[card.id] = card
		cards = append(cards, card)
	}
	
	// Vars for pseudo-recursion below
	totalCards := len(cards)
	newCards := cards

	// Pseudo-recursive method to get copies of cards, adding to total
	for len(newCards) > 0 {
		newCards = copies(newCards, cardsMap)
		totalCards += len(newCards)
	}
	
	return totalCards
}

func main() {
	// Read data from input file
	challengeData := utils.Input()

	// Print answers
	utils.Answers(strconv.Itoa(part1(challengeData)), strconv.Itoa(part2(challengeData)))
}