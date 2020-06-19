/*
if 和 else
在if的简短语句中声明的变量同样可以在任何对应的else
块中使用。

(在main 的 fmt.Println调用开始前，两次对pow的调用
均已经执行并返回其各自的结果。)
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64  {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main()  {
	fmt.Println(
		pow(3,2,10),
		pow(3,3,10),
	)
}
