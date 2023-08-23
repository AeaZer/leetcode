package linkedlist

// 61. 旋转链表
// https://leetcode.cn/problems/rotate-list/

/*给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	circleNode := head
	nodeLen := 1
	for circleNode.Next != nil {
		nodeLen++
		circleNode = circleNode.Next
	}
	circleNode.Next = head

	mod := nodeLen - (k % nodeLen)
	for i := 0; i < mod; i++ {
		circleNode = circleNode.Next
	}
	next := circleNode.Next
	circleNode.Next = nil
	circleNode = next

	return circleNode
}
