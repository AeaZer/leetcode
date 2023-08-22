package linkedlist

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// https://leetcode.cn/problems/merge-two-sorted-lists/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := new(ListNode)
	dummy := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			dummy.Next = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		} else {
			dummy.Next = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		}
		dummy = dummy.Next
	}
	if list1 != nil {
		dummy.Next = list1
	}
	if list2 != nil {
		dummy.Next = list2
	}
	return head.Next
}
