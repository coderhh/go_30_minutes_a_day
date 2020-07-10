package main

import "fmt"

type person struct {
	name string
	age  int
}

type T struct {
	ls []int
}

func foo(t T) {
	t.ls[0] = 100
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// fmt.Println(person{"Bob", 20})

	// fmt.Println(person{name: "Alice", age: 30})

	// fmt.Println(person{name: "Fred"})

	// fmt.Println(&person{name: "Ann", age: 40})

	// fmt.Println(newPerson("Jon"))

	// s := person{name: "Sean", age: 50}
	// fmt.Println(s.name)

	// sp := &s
	// fmt.Println(sp.age)

	// sp.age = 51
	// fmt.Println(sp.age)

	var t = T{
		ls: []int{1, 2, 3},
	}
	foo(t)
	fmt.Println(t.ls[0])

	isMatch := func(i int) bool {
		switch i {
		case 1:
		case 2:
			return true
		}
		return false
	}

	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}
