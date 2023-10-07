package heap

import "testing"

func TestKSmallestPairs(t *testing.T) {
	res := kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3)
	t.Log(res)
}
