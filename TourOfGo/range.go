/*
Range 
for 循环的 range 形式可以遍历切片或者映射。
当使用for循环遍历切片时，每次迭代都会返回两个值。
当第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
*/

package main 

import "fmt"

var pow  = []int{1,2,4,8,16,32,64,128}
func main()  {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}