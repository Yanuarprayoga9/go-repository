package main 

import (
	"fmt"
)

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width,height float64
}
func (r Rect) Area() float64 {
	return r.width * r.height
}
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main () {

	fmt.Println("interface ")

	var w Geometry = Rect{3, 4};
	fmt.Println(w.Area())
	fmt.Println(w.Perimeter());
}