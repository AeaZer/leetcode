package doublepointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	assert := assert.New(t)

	ok := isSubsequence("abc", "ahbgdc")
	assert.True(ok)

	ok = isSubsequence("axc", "ahbgdc")
	assert.False(ok)
}
