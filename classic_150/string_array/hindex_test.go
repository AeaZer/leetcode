package string_array

import "testing"

func TestHIndex(t *testing.T) {
	index := hIndex([]int{3, 0, 6, 1, 5})
	t.Log(index)
}
