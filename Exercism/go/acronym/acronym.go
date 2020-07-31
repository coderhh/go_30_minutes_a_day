// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) (out string) {
	a := regexp.MustCompile("[\\-\\_\\s]+")
	splitedString := a.Split(s, -1)
	//acronym := []string{}
	for _, word := range splitedString {
		out += string(word[0])
		// first := word[0:1]
		// acronym = append(acronym, first)
	}
	//acronymS := strings.Join(acronym, "")
	return strings.ToUpper(out)
}
