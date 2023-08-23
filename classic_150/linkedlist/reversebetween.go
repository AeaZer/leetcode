package linkedlist

/*给你单链表的头指针 head 和两个整数 left 和 right ，
其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。*/

// https://leetcode.cn/problems/reverse-linked-list-ii/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	pre := dummyHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyHead.Next
}
