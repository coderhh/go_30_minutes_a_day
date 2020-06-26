/*
接口与隐式实现


类型通过实现一个接口的所有方法来实现该接口。既然无需专门显示声明，也就没有 “implements”。
隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。

因此，也就无需在每一个实现上增加新的接口名称， 这样同时也鼓励了明确的接口定义。

*/
package main

import "fmt"
// I : this is an interface
type I interface{
	M()
}

// T : struct
type T struct {
	S string
}

// M  : 此方法表示类型T实现了接口I， 但我们无需显式声明此事。
func (t T) M()  {
	fmt.Println(t.S)
}

func main()  {
	var i I = T{"hello"}
	i.M()
}