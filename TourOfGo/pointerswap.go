package main

import (
	"fmt"
)

func swap(a,b *int) (*int, *int)  {
	a,b = b, a
	return a, b
}

func main()  {
	a, b : = 3, 4
	
	c, d := swap(&a,&b)

	a = *c
	b = *d

	fmt.Println(a, b)
}
