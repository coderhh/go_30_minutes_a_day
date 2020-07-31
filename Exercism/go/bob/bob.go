// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {

	var response string

	switch {
	case isQuestion(remark):
		response = "Sure."
	case isYellAt(remark):
		response = "Whoa, chill out!"
	case isYellAtQuestion(remark):
		response = "Calm down, I know what I'm doing!"
	case isSilence(remark):
		response = "Fine. Be that way!"
	default:
		response = "Whatever."

	}
	return response
}

func isQuestion(remark string) bool {
	match, _ := regexp.MatchString("^[A-Za-z0-9\\s:,!.)]*\\?(\\s)*$", remark)
	return match
}

func isYellAt(remark string) bool {

	match, _ := regexp.MatchString("^[^a-z?]*$", remark)

	match2, _ := regexp.MatchString("^[^a-zA-Z]*$", remark)
	return (match && !match2)
}

func isYellAtQuestion(remark string) bool {
	match, _ := regexp.MatchString("^[^a-z]*\\?$", remark)
	return match
}

func isSilence(remark string) bool {
	match, _ := regexp.MatchString("^\\s*$", remark)
	return match
}
