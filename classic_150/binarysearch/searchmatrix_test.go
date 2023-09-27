package binarysearch

import "testing"

func TestSearchMatrix(t *testing.T) {
	ok := searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}}, 3)
	t.Log(ok)
}
