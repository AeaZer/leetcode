package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert := assert.New(t)

	valid := isValid("(())")
	assert.True(valid)
}
