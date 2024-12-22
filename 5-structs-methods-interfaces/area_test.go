package main

import (
	"testing"
)

func TestArea(t *testing.T) {

	checkArea := func(shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %2f want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(circle, 314.1592653589793)
	})

	// Table driven tests (the same result as above)
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangles", Rectangle{w: 12, h: 6}, 72.0},
		{"circles", Circle{r: 10}, 314.1592653589793},
	}

	for _, v := range areaTests {
		t.Run(v.name, func(t *testing.T) {
			checkArea(v.shape, v.want)
		})
	}
}
