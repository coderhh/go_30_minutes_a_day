/*
指针接收者

你可以为指针接收者声明方法。

这意味着对于某类型T，接收者可以用*T 的文法。（此外，T不能是像 *int的这样的指针。）

例如，这里为*Vertex 定义了Scale方法。

指针接收者的方法可以修改接收者指向的值（就像Scale在这做的）。由于方法经常需要修改它的接收者，指针接收者
比值接收者更常用。

试着移除第16行scale函数声明中的*，观察此程序的行为如何变化。

若使用值接收者，那么Scale 方法会对原始Vertex值的副本进行操作。（对于函数的其它参数也是如此）
Scale方法必须用指针接收者来更改main函数中声明的Vertex的值。
*/

package main

import (
	"fmt"
    "math"
)

type Vertex struct {
	X,Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main()  {
	v :=Vertex{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())
}