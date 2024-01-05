package everyday

import "math"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 2487. 从链表中移除节点
// https://leetcode.cn/problems/remove-nodes-from-linked-list/description

/*给你一个链表的头节点 head 。
移除每个右侧有一个更大数值的节点。
返回修改后链表的头节点 head 。*/

// head = [5,2,13,3,8]

func removeNodes(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  math.MaxInt,
		Next: head,
	}
	pre, slow, fast := dummyHead, head, head.Next
	for slow.Next != nil {
		for fast != nil && fast.Val <= slow.Val {
			fast = fast.Next
		}
		if fast != nil {
			pre.Next = slow.Next
		} else {
			pre = slow
		}
		slow = slow.Next
		if fast == nil || slow == fast {
			fast = slow.Next
		}
	}
	return dummyHead.Next
}
