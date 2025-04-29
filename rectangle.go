// отработка указтелей и методов
package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Scale(factor float64) {
	r.length = r.length * factor
	r.width = r.width * factor
}

func main() {
	rectangle := Rectangle{10, 12}

	fmt.Printf("Area: %f\n", rectangle.Area())
	fmt.Printf("Perimeter: %f\n", rectangle.Perimeter())
	rectangle.Scale(3.)
	fmt.Printf("Scale: %+v\n", rectangle)

}
