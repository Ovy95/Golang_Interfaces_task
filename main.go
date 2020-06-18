package main

import "fmt"

type shape interface {
	getArea() float64
}

type trangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	right := trangle{6.4, 10}
	square1 := square{5.0}

	getArea(right)
	getArea(square1)

}

func getArea(s shape) {
	fmt.Println(s.getArea())
}

func (t trangle) getArea() float64 {
	return 0.5 * t.base * t.height

}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
