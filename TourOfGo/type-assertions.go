/*
类型断言
类型断言提供了访问接口值底层具体值的方式。
t ：= i.(T)

该语句断言接口值i保存了具体类型T，并将其底层类型
为T的值赋予变量t。
若i 并未保存T类型的值， 该语句就会触发一个恐慌。

为了判断一个接口值是否存在了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

*/

package main

import "fmt"

func main()  {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s,ok)
	
	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错（panic)
	fmt.Println(f)
}