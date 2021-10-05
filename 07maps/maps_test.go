package _7pointersanderrors

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("i am able to search in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		actual, _ := dictionary.Search( "test")
		expected := "this is just a test"

		assert.Equal(t,actual,expected)
	})

	t.Run("i get an error message when word does not exist in dict", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_ , err := dictionary.Search( "hello")
		expected := ErrNotFound

		assert.Equal(t,err,expected)
	})

	t.Run("i am able to add ", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		dictionary.AddWord("ok","yes its ok")
		actual , _ := dictionary.Search( "ok")
		assert.Equal(t,actual,"yes its ok")
	})
}

