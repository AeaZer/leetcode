package breadthfirstsearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinMutation(t *testing.T) {
	assert := assert.New(t)

	i := minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AAACGGTA", "AACCGCTA"})
	assert.Equal(2, i)
}
