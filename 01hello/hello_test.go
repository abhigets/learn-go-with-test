package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("with empty argument",func(t *testing.T){
		actualValue := Hello("", "")
		expectedValue := "Hello, World!"
		assert.Equal(t,actualValue,expectedValue)
	})

	t.Run("with an valid argument",func(t *testing.T){
		actualValue := Hello("Chris", "")
		expectedValue := "Hello, Chris"
		assert.Equal(t,actualValue,expectedValue)
	})

	t.Run("in spanish",func(t *testing.T){
		actualValue := Hello("Elodia","Spanish")
		expectedValue := "Hola, Elodia"
		assert.Equal(t, actualValue, expectedValue)
	})

	t.Run("in french",func(t *testing.T){
		actualValue := Hello("Elodia","French")
		expectedValue := "Bonjour, Elodia"
		assert.Equal(t, actualValue, expectedValue)
	})
}