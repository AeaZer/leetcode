package breadthfirstsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMutation(t *testing.T) {
	assert := assert.New(t)

	i := minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AAACGGTA", "AACCGCTA"})
	assert.Equal(2, i)
}
