package main

import "fmt"

type area interface {
	calculateArea(side int)
}

type shape struct {
	geometry area
}
type circle struct{}

func (c circle) calculateArea(radius int) {
	fmt.Println("Area of circle is", (3.14 * float64(radius) * float64(radius)))
}

type square struct{}

func (s square) calculateArea(side int) {
	fmt.Println("Area of square is ", float64(side*side))
}

func (s shape) calculate(side int) {
	s.geometry.calculateArea(side)
}
func main() {
	circle := circle{}
	square := square{}

	circleshape := shape{geometry: circle}
	squareshape := shape{geometry: square}

	circleshape.calculate(5)
	squareshape.calculate(6)

}
