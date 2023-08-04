package string_array

import "testing"

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	rotate(nums, 2)

	t.Log(nums)
}
