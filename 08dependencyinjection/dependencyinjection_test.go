package main

import (
	"bytes"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("i am able to inject a dependency", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		actual := buffer.String()
		expected := "Hello, Chris"
		assert.Equal(t,actual,expected)
	})
}

