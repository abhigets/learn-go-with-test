package main

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {

	t.Run("i am able to mock", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySleeper := SpySleeper{}

		Countdown(&buffer, spySleeper)

		actual := buffer.String()
		expected := `3
2
1
Go!`

		assert.Equal(t, actual, expected)
		assert.Equal(t, spySleeper.Calls, 4)
	})
}
