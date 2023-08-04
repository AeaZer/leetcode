package string_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJump(t *testing.T) {
	assert := assert.New(t)

	ok := canJump([]int{2, 3, 1, 1, 4})
	assert.True(ok)

	ok = canJump([]int{2, 0, 0})
	assert.True(ok)

	ok = canJump([]int{3, 2, 1, 0, 4})
	assert.False(ok)
}
