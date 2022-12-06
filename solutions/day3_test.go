package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneToPriority(t *testing.T) {
	assert.Equal(t, 1, byteToPriority('a'))
	assert.Equal(t, 26, byteToPriority('z'))
	assert.Equal(t, 27, byteToPriority('A'))
	assert.Equal(t, 52, byteToPriority('Z'))
	assert.Equal(t, 0, byteToPriority('-'))
}

func TestIsInBag(t *testing.T) {
	var bag string = "ABCDE"
	assert.True(t, isInBag(65, bag))
	assert.False(t, isInBag(100, bag))
}
