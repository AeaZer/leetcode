package string_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	assert := assert.New(t)

	answer := productExceptSelf([]int{1, 2, 3, 4})
	assert.Equal([]int{24, 12, 8, 6}, answer)
}
