package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	e := 0.00000001
	for math.Abs((math.Pow(z,2) - x) / (2 * z)) > e{
		z -= (math.Pow(z,2) - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}