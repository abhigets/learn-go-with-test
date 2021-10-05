package _3iteration

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestRepeat(t *testing.T){
	t.Run("test Repeat", func(t *testing.T) {
		actual := Repeat("a",6)
		expected := "aaaaaa"

		assert.Equal(t, actual,expected)
	})

	t.Run("test builder", func(t *testing.T) {
		actual := Builder()
		expected := "1...2...3...repeat"
		assert.Equal(t, actual,expected)
	})


}

func BenchmarkRepeat(b *testing.B) {
	for i:=0 ; i < b.N ; i++{
		Repeat("a", 6)
	}
}

