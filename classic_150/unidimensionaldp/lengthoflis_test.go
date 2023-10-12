package unidimensionaldp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	lis := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	t.Log(lis)
}
