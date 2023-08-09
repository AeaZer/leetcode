package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := assert.New(t)

	res := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(3, res)

	res = lengthOfLongestSubstring("bbbbbbb")
	assert.Equal(1, res)

	res = lengthOfLongestSubstring("pwwkew")
	assert.Equal(3, res)
}
