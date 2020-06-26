/*
类型选择

类型选择是一种按顺序从几个类型断言中选择分支的结构。

类型选择与一般的switch语句相似，不过类型选择中的case为类型（而非值），它们针对给定接口值
所存储的值的类型进行比较。

switch v := i.(type)
case T:
	// v 的类型为 T
case S:
	// v 的类型为 S
default:
	// 没有匹配，v 与 i 的类型相同

类型选择中的声明与类型断言i.(T)的语法相同，只是具体类型T被替换成了关键字 type。

此选择语句判断接口值i保存的值类型是T还是S。 在T或S的情况下，变量v会分别按T或S类型保存i
拥有的值。 在默认（即没有匹配）的情况下， 变量v与i的接口类型和值相同。
*/

package main

import "fmt"

func do(i interface{})  {
	switch v := i.(type){
	case int:
		fmt.Printf("Twice %v is %v \n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main()  {
	do(21)
	do("hello")
	do(true)
}