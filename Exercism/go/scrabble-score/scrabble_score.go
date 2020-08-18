package scrabble

import (
	"strings"
)

const ones string = "AEIOULNRST"
const twos string = "DG"
const threes string = "BCMP"
const fours string = "FHVWY"
const fives string = "K"
const eights string = "JX"
const tens string = "QZ"

// Score  Score method
func Score(input string) int {
	score := 0
	input = strings.ToUpper(input)

	for i := range input {
		switch {
		case strings.Contains(ones, string(input[i])):
			score = score + 1
		case strings.Contains(twos, string(input[i])):
			score = score + 2
		case strings.Contains(threes, string(input[i])):
			score = score + 3
		case strings.Contains(fours, string(input[i])):
			score = score + 4
		case strings.Contains(fives, string(input[i])):
			score = score + 5
		case strings.Contains(eights, string(input[i])):
			score = score + 8
		case strings.Contains(tens, string(input[i])):
			score = score + 10
		default:

		}

	}
	return score
}
