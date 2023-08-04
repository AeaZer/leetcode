package string_array

import "testing"

func TestJump(t *testing.T) {
	minThey := jump([]int{2, 1, 1, 1, 1})
	t.Log(minThey)
}
