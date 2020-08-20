package reverse

import "bytes"

// Reverse method return the reverse of a string
func Reverse(input string) string {

	// var newString string = ""

	// for i := len(input) - 1; i >= 0; i-- {
	// 	newString += string(input[i])
	// }

	// return newString

	var output bytes.Buffer

	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		output.WriteRune(runes[i])
	}

	return output.String()

}
