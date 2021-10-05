package _4arary

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestArray(t *testing.T) {
	t.Run("Sum of array", func(t *testing.T) {
		number := []int{1,2,3,4,5}

		actual := Sum(number)
		expect := 15

		assert.Equal(t,actual,expect)
	})

	t.Run("Sum all the arrays", func(t *testing.T) {
		array1 := []int{1,2,3,4,5}
		array2 := []int{1,2,3,4,5,6}

		actual := SumAll(array1,array2)
		expect := []int{15,21}

		assert.Equal(t,actual,expect)
	})

	t.Run("Sum all tails", func(t *testing.T) {
		array1 := []int{1,2}
		array2 := []int{0,9}

		actual := SumOfTails(array1,array2)
		expect := []int{2,9}

		assert.Equal(t,actual,expect)
	})

	t.Run("Sum all tails with empty array", func(t *testing.T) {
		var array1 []int
		array2 := []int{0,9}

		actual := SumOfTails(array1,array2)
		expect := []int{0,9}

		assert.Equal(t,actual,expect)
	})
}

