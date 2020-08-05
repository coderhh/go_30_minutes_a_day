package hamming

import (
	"errors"
)

// Distance method
func Distance(a, b string) (int, error) {

	distance := 0
	err := errors.New("error")
	if len(a) != len(b) {
		distance = 0
		err = errors.New("can't work")
	} else {
		for i := range a {
			if a[i] != b[i] {
				distance = distance + 1
			}
		}
		err = nil
	}
	return distance, err
}
