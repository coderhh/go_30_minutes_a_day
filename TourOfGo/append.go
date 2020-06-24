/*
向切片追加元素

为切片追加新的元素是一种常用的操作，为此Go提供了内建的
append函数。内建函数的文档对此函数有详细的介绍。
func append(s []T, vs ...T) []T

append 函数的第一个参数s 是一个元素类型为T的切片， 其余类型
为T的值将会追加到该切片的末尾。
append 的结果是每一个包含原切片所有元素加上新添加的元素的切片。
当s 的底层数组太小，不足以容纳所有给定的值时， 它就会分配一个更大
数组。返回的切片会指向这个新分配的数组。

*/

package main

import "fmt"

func main()  {
	var s []int
	printSlice(s)

	//添加一个空切片
	s = append(s, 0)
	printSlice(s)

	//这个切片会按需增长
	s = append(s, 1)
	printSlice(s)
	//可以一次性添加多个元素
	s = append(s, 2,3,5)
	printSlice(s)
	
}

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s),cap(s),s)
}