package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	assert := assert.New(t)

	res := minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	assert.Equal(2, res)

	res = minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})
	assert.Equal(0, res)

}
