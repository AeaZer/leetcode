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

func TestRemoveNodes(t *testing.T) {
	// [5,2,13,3,8]
	node := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}
	removeNodes(node)
}
