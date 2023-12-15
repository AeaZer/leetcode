package everyday

import "testing"

func TestHIndex(t *testing.T) {
	hIndex([]int{3, 0, 6, 1, 5})
}

func TestCountPoints(t *testing.T) {
	countPoints("B0B6G0R6R0R6G9")
}

func TestReverseOddLevels(t *testing.T) {
	node := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 13,
		},
		Right: &TreeNode{
			Val: 11,
		},
	}
	reverseOddLevels(node)
}
