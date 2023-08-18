package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
	assert := assert.New(t)

	stack := Constructor()
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)
	assert.Equal(1, stack.GetMin())
	stack.Pop()
	assert.Equal(2, stack.GetMin())
	top := stack.Top()
	assert.Equal(2, top)
}
