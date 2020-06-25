package main

import "fmt"

func fibonacci() func() int  {
	a, b := 0, 1
	if a == 0 {
		fmt.Println(a)
	}

	return func () int  {
		a, b = b, a + b
		return a 
		
	}
	
}

func main()  {
	f := fibonacci()
	for i :=0; i < 10; i++ {
		fmt.Println(f())
	}
}