package doublepointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert := assert.New(t)

	ok := isPalindrome("A man, a plan, a canal: Panama")
	assert.True(ok)

	ok = isPalindrome("race a car")
	assert.False(ok)
}
