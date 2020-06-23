/*
nil 切片

切片的零值是nil。
nil 切片的长度和容量为0且没有底层数组。

*/

package main

import "fmt"

func main()  {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}