package raindrops

import "strconv"

// Convert method
func Convert(input int) (result string) {

	if (input%3 != 0) && (input%5 != 0) && (input%7 != 0) {
		result = strconv.Itoa(input)
	}

	if input%3 == 0 {
		result = result + "Pling"
	}
	if input%5 == 0 {
		result = result + "Plang"
	}

	if input%7 == 0 {
		result = result + "Plong"
	}

	return result
}
