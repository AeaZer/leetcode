package string_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCompleteCircuit(t *testing.T) {
	assert := assert.New(t)

	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	completeCircuit := canCompleteCircuit(gas, cost)
	assert.Equal(3, completeCircuit)
}
