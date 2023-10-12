package unidimensionaldp

import "testing"

func TestCoinChange(t *testing.T) {
	change := coinChange([]int{1, 2, 5}, 11)
	t.Log(change)
}
