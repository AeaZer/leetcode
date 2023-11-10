package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func createLinkedList(nums []int) *ListNode {
	head := new(ListNode)
	dummyHead := head
	for _, num := range nums {
		dummyHead.Next = &ListNode{
			Val: num,
		}
		dummyHead = dummyHead.Next
	}
	return head.Next
}

func TestReverseList(t *testing.T) {
	require := require.New(t)

	node := createLinkedList([]int{1, 2, 3, 4, 5})
	require.NotNil(node)

	reverseList(node)
}
