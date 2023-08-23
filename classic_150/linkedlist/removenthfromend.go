package linkedlist

/*给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。*/

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmpNode := head
	var l int
	for tmpNode != nil {
		l++
		tmpNode = tmpNode.Next
	}

	dummyNode := new(ListNode)
	dummyNode.Next = head
	pre := dummyNode
	n = l - n
	for i := 0; i < n; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	pre.Next = cur.Next
	return dummyNode.Next
}
