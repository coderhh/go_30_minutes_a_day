/*
没有条件的switch

没有条件的switch同switch true 一样。
这种形式能将一长串if-then-else 写得更加清楚
*/

package main 

import (
	"fmt"
	"time"
)

func main()  {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}