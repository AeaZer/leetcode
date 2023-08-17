package interval

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	assert := assert.New(t)

	strSlice := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	assert.Equal([]string{"0->2", "4->5", "7"}, strSlice)
}
