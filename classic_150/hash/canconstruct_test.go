package hash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	assert := assert.New(t)

	ok := canConstruct("aa", "aab")
	assert.True(ok)
}
