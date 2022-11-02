package main

// 61. 旋转链表
// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func newListNode() *ListNode {
	listNode := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	return &listNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	loopNode := head
	for loopNode.Next != nil {
		loopNode = loopNode.Next
		n++
	}
	mod := n - k%n
	if mod == n {
		return head
	}
	loopNode.Next = head
	for mod > 0 {
		loopNode = loopNode.Next
		mod--
	}
	ret := loopNode.Next
	loopNode.Next = nil
	return ret
}
