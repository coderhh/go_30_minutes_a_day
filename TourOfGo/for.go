/*
for

Go 只有一种循环结构：for 循环。
基本 for 循环由三部分组成，它们用分号隔开：
	初始化语句： 在第一次迭代前执行
	条件表达式： 在每次迭代前求值
	后置语句： 在每次迭代的结尾执行
初始化语句通常为一句短变量声明，该变量声明仅在for语句的
作用域中可见。

一旦条件表达式布尔值为false,循环迭代就会终止。

注意： 和 C java javascript 之类的语言不同，Go 的 for 语句后面
的三个构成部分外没有小括号，大括号 {} 则是必须的。
*/

package main

import "fmt"

func main() {
	sum :=  0
	for i :=0; i < 10; i++ {
		sum +=i
	}
	fmt.Println(sum)
}