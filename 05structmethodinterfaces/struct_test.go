package _5structmethodinterfaces

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func PerimeterUsingStruct(t *testing.T) {
	t.Run("perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		actual := Perimeter(rectangle)
		expected := 40.0

		assert.Equal(t, actual, expected)
	})
}

func AreaUsingStruct(t *testing.T) {
	t.Run("area for rectangle", func(t *testing.T) {
		rectangle := Rectangle{6.0, 7.0}

		actual := rectangle.Area()
		expected := 42.0

		assert.Equal(t, actual, expected)
	})

	t.Run("area of circle", func(t *testing.T) {
		circle := Circle{10.0}

		actual := circle.Area()
		expected := 314.1

		assert.Equal(t, actual, expected)
	})
}

func TestAreaWithStruct(t *testing.T) {
	t.Run("test the table driven test", func(t *testing.T) {
		areaTest := []struct {
			shape    Shape
			expected float64
		}{
			{Rectangle{Width: 10, Height: 20}, 200.0},
			{Circle{Radius: 10}, 314.1592653589793},
			{Triangle{Length: 12, Height: 6}, 36},
		}

		for _, tt := range areaTest {
			actual := tt.shape.Area()
			expected := tt.expected
			assert.Equal(t, actual, expected)
		}
	})

}
