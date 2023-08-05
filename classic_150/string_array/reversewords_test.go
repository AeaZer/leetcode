package string_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	assert := assert.New(t)

	words := reverseWords(" asdasd df f")
	assert.Equal("f df asdasd", words)

	words = reverseWords("the sky is blue")
	assert.Equal("blue is sky the", words)

	words = reverseWords("  hello world  ")
	assert.Equal("world hello", words)

	words = reverseWords("a good   example")
	assert.Equal("example good a", words)

}
