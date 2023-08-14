package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	assert := assert.New(t)

	ok := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)
	assert.True(ok)

	ok = containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2)
	assert.False(ok)
}
