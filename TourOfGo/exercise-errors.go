package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt : New error type
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number :  %v\n", float64(e))
}

// Sqrt : 
func Sqrt(x float64) (float64, error) {
	if x > 0 {
		z := float64(1)
		e := 0.00000001
		for math.Abs((math.Pow(z, 2)-x)/(2*z)) > e {
			z -= (math.Pow(z, 2) - x) / (2 * z)
		}
		return z, nil
	} 
	return 0, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}