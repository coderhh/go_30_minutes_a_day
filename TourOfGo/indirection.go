/*
方法与指针重定向

比较前两个程序， 你大概会注意到带指针参数的函数必须接受一个指针：
var v Vertex

ScaleFunc(v, 5) //编译错误！
ScaleFunc(&v, 5) //OK

而以指针为接收者的方法被调用时， 接收者既能为值又能为指针：
var v Vertext
v.Scale(5) // OK
p :=&v
p.Scale(10) //OK
对于语句v.Scale(5), 即便v 是个值而非指针， 带指针接收者的方法也能被直接调用。也就是说，
由于Scale方法有一个指针接收者，为方便起见，Go会将语句v.Scale(5)解释为（&v）.Scale(5)。
*/

package main

import "fmt"

type Vertex struct {
	X,Y float64
}

func (v *Vertex) Scale(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main()  {
	v := Vertex{3,4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4,3}
	p.Scale(3)
	ScaleFunc(p,8)

	fmt.Println(v, p)
}