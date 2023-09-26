package dividerule

import "sort"

// 148. 排序链表
// https://leetcode.cn/problems/sort-list/?envType=study-plan-v2&envId=top-interview-150
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

func sortList(head *ListNode) *ListNode {
	elems := make([]int, 0)
	for head != nil {
		elems = append(elems, head.Val)
		head = head.Next
	}
	sort.Ints(elems)
	newNode := new(ListNode)
	dummyNode := newNode
	for i := range elems {
		dummyNode.Next = &ListNode{
			Val: elems[i],
		}
		dummyNode = dummyNode.Next
	}
	return newNode.Next
}
