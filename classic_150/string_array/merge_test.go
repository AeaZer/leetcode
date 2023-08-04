package string_array

import "testing"

func TestMerge(t *testing.T) {

	nums1 := []int{4, 0, 0, 0, 0, 0}

	merge(nums1, 1, []int{1, 2, 3, 5, 6}, 5)
	t.Log(nums1)
}
