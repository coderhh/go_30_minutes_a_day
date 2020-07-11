package main

import "fmt"

type rect struct {
	width, height, area, perim int
}

func (r *rect) areaFun() {
	r.area = r.width * r.height
}

func (r rect) perimFun() {
	r.perim = 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5, area: 0, perim: 0}
	r.areaFun()
	r.perimFun()
	fmt.Println("area: ", r.area)
	fmt.Println("perim: ", r.perim)

	rp := &r

	fmt.Println("area: ", rp.area)
	fmt.Println("perim: ", rp.perim)
}
