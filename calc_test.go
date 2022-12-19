package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusSimply(t *testing.T) {
	result := Plus(1, 3)
	expect := 4

	assert.Equal(t, expect, result)
}

