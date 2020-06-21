/*
指针
Go拥有指针。指针保存了值的内存地址。
类型*T 是指向T类型值的指针。其值为nil。
var p *int
&操作符会生成一个指向操作数的指针。

i := 42
p = &i

* 操作符表示指针指向的底层值。

fmt.Println(*p) // 通过指针p取值i
*p = 21         // 通过指针p 设置 i

这也就是通常所说的“简介引用”或“重定向”。
与 C 不同， Go没有指针运算。
*/

package main

import "fmt"

func main()  {
	i, j := 42, 2701

	p := &i          // 指向i
	fmt.Println(*p)  // 通过指针读取 i 的值
	*p = 21          // 通过指针设置 i 的值
	fmt.Println(i)

	p = &j           // 指向 j
	*p = *p / 37     // 通过指针对 j 进行除法运算
	fmt.Println(j)   // 查看 j 的值
	
}