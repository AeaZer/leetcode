package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNode_ValueInBST(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewBSTWithSlice(nums)

	inBST := tree.ValueInBST(3)
	assert.True(inBST)

	inBST = tree.ValueInBST(6)
	assert.True(inBST)

	inBST = tree.ValueInBST(8)
	assert.False(inBST)
}

func TestNode_InsertValue(t *testing.T) {
	nums := []int{1, 2, 4, 5, 6, 7, 8}
	tree := NewBSTWithSlice(nums)
	value := tree.InsertValue(4)
	slice := RangeBSTToSlice(value)
	t.Log(slice)
}

func TestNode_SearchKValue(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewBSTWithSlice(nums)

	value, ok := tree.SearchKValue(2)
	require.True(ok)
	assert.Equal(2, value)

	value, ok = tree.SearchKValue(7)
	require.True(ok)
	assert.Equal(7, value)
}

func TestNewBSTWithSlice(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewBSTWithSlice(nums)
	require.NotNil(tree)
	assert.EqualValues(4, tree.value)
	assert.EqualValues(1, tree.left.left.value)
	assert.EqualValues(2, tree.left.value)
	assert.EqualValues(3, tree.left.right.value)
	assert.EqualValues(5, tree.right.left.value)
	assert.EqualValues(6, tree.right.value)
	assert.EqualValues(7, tree.right.right.value)
}

func TestRangeBSTToSlice(t *testing.T) {
	assert := assert.New(t)

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	tree := NewBSTWithSlice(nums)
	r := RangeBSTToSlice(tree)
	assert.Equal(nums, r)
}
