package string_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedSet(t *testing.T) {
	assert := assert.New(t)

	r := Constructor()
	ok := r.Insert(1)
	assert.True(ok)
	assert.Equal([]int{1}, r.vales)

	ok = r.Insert(1)
	assert.False(ok)
	assert.Equal([]int{1}, r.vales)

	ok = r.Insert(2)
	assert.True(ok)
	assert.Equal([]int{1, 2}, r.vales)

	ok = r.Insert(4)
	assert.True(ok)
	assert.Equal([]int{1, 2, 4}, r.vales)

	ok = r.Insert(3)
	assert.True(ok)
	assert.Equal([]int{1, 2, 3, 4}, r.vales)

	ok = r.Remove(3)
	assert.True(ok)
	assert.Equal([]int{1, 2, 4}, r.vales)

	ok = r.Remove(6)
	assert.False(ok)
	assert.Equal([]int{1, 2, 4}, r.vales)

	value := r.GetRandom()
	t.Log(value)
	assert.NotZero(value)
}
