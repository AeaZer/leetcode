package binarysearch

import "testing"

func TestSearchInsert(t *testing.T) {
	index := searchInsert([]int{1, 2, 5, 6}, 2)
	t.Log(index)
}
