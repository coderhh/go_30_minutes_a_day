package isogram

import (
	"regexp"
	"strings"
)

// IsIsogram return true if the input is an isogram
func IsIsogram(input string) bool {

	s := make([]string, 0)

	r, _ := regexp.Compile("[^a-zA-Z]")
	input = r.ReplaceAllString(input, "")
	input = strings.ToUpper(input)
	for i := range input {
		_, found := Find(s, string(input[i]))
		if found {
			return false
		}
		s = append(s, string(input[i]))
	}

	return true
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
