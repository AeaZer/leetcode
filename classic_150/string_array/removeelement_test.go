package string_array

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{1}
	removeElement(nums, 1)

	t.Log(nums)
}

func TestName(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	copy(nums[0:], nums[1:])
	t.Log(nums)
}
