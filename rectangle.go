// отработка указтелей и методов
package main

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
