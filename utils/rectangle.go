// отработка указтелей и методов
package utils

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (r *Rectangle) Scale(factor float64) {
	r.Length = r.Length * factor
	r.Width = r.Width * factor
}
