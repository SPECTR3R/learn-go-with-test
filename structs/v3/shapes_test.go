package shapes

import "testing"

func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 15, Height: 6, a: 9, b: 9}, hasArea: 45.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTest := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPerimeter: 36.0},
		{name: "Circle", shape: Circle{Radius: 5}, hasPerimeter: 31.41592653589793},
		{name: "Triangle", shape: Triangle{Base: 15, Height: 6, a: 9, b: 9}, hasPerimeter: 33.0},
	}

	for _, tt := range perimeterTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}
