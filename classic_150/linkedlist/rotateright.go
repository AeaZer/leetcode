package linkedlist

// 61. 旋转链表
// https://leetcode.cn/problems/rotate-list/

/*给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。*/

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 创建循环链表并计算初链表的长度
	circleNode := head
	nodeLen := 1
	for circleNode.Next != nil {
		nodeLen++
		circleNode = circleNode.Next
	}
	circleNode.Next = head

	// 取模操作，我们发现每移动 nodeLen 次，移动后的链表和原链表相同
	mod := nodeLen - (k % nodeLen)
	for i := 0; i < mod; i++ {
		circleNode = circleNode.Next
	}
	// 断开循环链表
	next := circleNode.Next
	circleNode.Next = nil
	circleNode = next

	return circleNode
}
