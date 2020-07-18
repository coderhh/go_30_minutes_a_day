package main

import (
	"fmt"
)

func swap(a,b *int) (*int, *int)  {
	
	fmt.Println(a,b)
	a,b = b, a
	
	fmt.Println(a,b)
	return a, b
}

func main()  {
	a, b := 3, 4

	a,b = b,a

	fmt.Println(a, b)
	
	c, d := swap(&a,&b)

	fmt.Println(*c, *d)
	a = *c
	b = *d
	
	fmt.Println(a, b)
}