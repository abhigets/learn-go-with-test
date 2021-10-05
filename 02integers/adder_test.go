package _2integers

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAdder(t *testing.T){
	sum := Add(2,2)
	expected := 4

	assert.Equal(t, sum, expected)
}

