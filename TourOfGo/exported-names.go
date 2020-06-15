/*
导出名

在GO中，如果一个名字以大写字母开头，那么它就会已经导出的。例如，Pizza 
就是个已经导出名， Pi也同样， 它导出自 math 包。

pizza 和 pi 并未以大写字母开头，所以它们是未导出的。

在导入一个包时， 你只能引用其中已经导出的名字。任何“未导出”的名字在该
包外均无法访问。

执行代码， 观察错误输出。

然后将 math.pi 改名为 math.Pi 再试着执行一次。
*/



package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(math.Pi)
}