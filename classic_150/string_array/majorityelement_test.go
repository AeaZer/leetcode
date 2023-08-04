package string_array

import "testing"

func TestMajorityElement(t *testing.T) {
	element := majorityElement([]int{1, 1, 1, 2, 2})
	t.Log(element)
}
