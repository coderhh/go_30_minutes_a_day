/*
switch

switch 是编写一连串if - else 语句的简便方法。它运行第一个值
等于表达式的case语句。
GO的switch语句类似于C，C++，Java，javascrip和PHP中，不过Go只运行选定
的case，而非之后所有的case。实际上，Go自动提供了在这些语句中每个case后面
所需的break语句。除非以fallfthrough语句结束，否则分支会自动终止。Go的另一点
重要的不同在于switch的case无需为常量，且取值不必为整数。
*/

package main

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}