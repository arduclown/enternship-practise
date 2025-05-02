// запуск через go test -v

package main

import (
	"testing"

	"github.com/arduclown/enternship-practise/utils"
)

func TestArea(t *testing.T) {
	r := utils.Rectangle{Length: 5, Width: 3}
	excepted := 15.
	if got := r.Area(); got != excepted {
		t.Errorf("Area() = %v, want %v", got, excepted)
	}

}

func TestPerimeter(t *testing.T) {
	r := utils.Rectangle{Length: 5, Width: 3}
	excepted := 16.
	if got := r.Perimeter(); got != excepted {
		t.Errorf("Perimeter() = %v, want %v", got, excepted)
	}
}

func TestScale(t *testing.T) {
	r := utils.Rectangle{Length: 5, Width: 3}
	excepted1 := 25.
	excepted2 := 15.
	r.Scale(5.)
	if r.Length != excepted1 || r.Width != excepted2 {
		t.Errorf("Rectangle = %v, want %v, %v", r, excepted1, excepted2)
	}
}
