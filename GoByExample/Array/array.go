package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{6, 2, 1, 3, 5}
	fmt.Println("dcl:", b)

	bubbleSort(&b)
	fmt.Println("sorted_dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func bubbleSort(array *[5]int) {
	len := len(*array)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if (*array)[j] > (*array)[j+1] {
				temp := (*array)[j]
				(*array)[j] = (*array)[j+1]
				(*array)[j+1] = temp
			}
		}
	}
}
