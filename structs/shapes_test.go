package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	actual := Perimeter(Rectangle{10.0, 10.0})
	expected := 40.0
	if expected != actual {
		t.Errorf("got %.2f want %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expected: 136},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			if actual != tt.expected {
				t.Errorf("%#v actual %g want %g", tt.shape, actual, tt.expected)
			}
		})
	}
}
