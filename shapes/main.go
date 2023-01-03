package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{sideLength: 2.0}
	tr := triangle{base: 2.0, height: 3.0}
	printArea(sq)
	printArea(tr)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println("Area of shape", s, "is", s.getArea())
}
