/*
指针与函数
我们要把Abs 和 Scale 方法充血为函数。
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64  {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main()  {
	v := Vertex{3,4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}