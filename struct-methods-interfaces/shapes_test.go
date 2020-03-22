package main

import "testing"

func TestPerimeter(t *testing.T) {
	want := 40.0

	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rectangle)

	if got != want {
		t.Errorf("want %.2f, got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("want %g, got %g", want, got)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	want := 72.0
	// 	rectangle := Rectangle{Width: 12.0, Height: 6.0}

	// 	checkArea(t, rectangle, want)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	want := 314.1592653589793
	// 	circle := Circle{Radius: 10.0}

	// 	checkArea(t, circle, want)
	// })

	areaTests := []struct {
		name       string
		shape      Shape
		wantedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, wantedArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, wantedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, wantedArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.wantedArea {
				t.Errorf("%#v got an Area of %g wanted an Area of %g", tt.shape, got, tt.wantedArea)
			}
		})

	}
}
