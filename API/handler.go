//go:build file2
// +build file2

package main

import "fmt"

const pi float64 = 3.14

type Rectangle struct {
	length, breadth int
}

type Circle struct {
	radius int
}

/*
value handlers/receivers dont modify the values of the structure it is binded to.
Does the operations using a COPY. (Pass by Value).
*/
func (r Rectangle) getArea() (area int) {
	r.length, r.breadth = 20, 10
	area = r.length * r.breadth
	return area
}

/*
reference handlers/receivers  modify the values of the structure it is binded to.
Does the operations using a REFERENCE. (Pass by REFERENCE).
*/
func (r *Rectangle) getTripledArea() (area int) {
	r.length, r.breadth = r.length*3, r.breadth*3
	area = r.length * r.breadth
	return area
}

func (c Circle) getCircleArea() (area float64) {
	area = float64(c.radius) * pi
	return area
}

func main() {
	var rect Rectangle = Rectangle{length: 10, breadth: 5}
	fmt.Printf("Area of the rect is: %d", rect.getArea())
	fmt.Printf("\nLength: %d Breadth: %d", rect.length, rect.breadth)

	fmt.Printf("\nDoubled area of the rect is: %d", rect.getTripledArea())
	fmt.Printf("\nLength: %d Breadth: %d", rect.length, rect.breadth)

	var c Circle = Circle{radius: 10}
	fmt.Printf("\nArea of the Circle having radius %d : %f", c.radius, c.getCircleArea())
}
