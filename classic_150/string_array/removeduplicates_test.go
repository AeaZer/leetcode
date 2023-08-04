package string_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	n := removeDuplicates(nums)
	t.Log(n)
}
