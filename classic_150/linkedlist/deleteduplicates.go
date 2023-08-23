package linkedlist

/*给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。*/

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/

// NULL [1,2,3,3,4,5]

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -101}
	dummyHead.Next = head
	cur := dummyHead

	for cur.Next != nil && cur.Next.Next != nil {
		// 比较后继节点和后继的后继节点是否相同，不相同就进入下一个节点判断
		if cur.Next.Val != cur.Next.Next.Val {
			cur = cur.Next
			continue
		}
		// 相等的话就切掉相等的部分
		cur.Next = unEqualHeadNode(cur.Next)
	}
	return dummyHead.Next
}

func unEqualHeadNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	val := head.Val
	dummy := head.Next
	for dummy != nil && dummy.Val == val {
		dummy = dummy.Next
	}
	return dummy
}
