package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64  {
	z := float64(0.5)
	e := 0.00000001
	for {
		y := (math.Pow(z,2) - x) / (2 * z)
		if y < 0{
			y = -y
		}
		if y < e{
			break
		}
		z -= (math.Pow(z,2) - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main()  {
	fmt.Println(Sqrt(2))
}