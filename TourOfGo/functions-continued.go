/*
函数（续）

当连续两个或者多个函数的已命名形参类型相同时，除最后一个类型
以外，其他都可以省略。

在本例中，

 x int, y int

被缩写为
 x, y int
*/


package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main(){
	fmt.Println(add(42, 13))
}