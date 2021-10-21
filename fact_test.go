package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	j = factJson{
		Text: "fact",
	}
)

func TestString(t *testing.T) {
	assert.Equal(t, j.String(), "fact")
}
