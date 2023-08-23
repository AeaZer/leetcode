package linkedlist

// 86. 分隔链表
// https://leetcode.cn/problems/partition-list/

/*给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。*/

func partition(head *ListNode, x int) *ListNode {
	left, right := new(ListNode), new(ListNode)
	dummyLeft, dummyRight := left, right
	for head != nil {
		if head.Val < x {
			dummyLeft.Next = head
			dummyLeft = dummyLeft.Next
		} else {
			dummyRight.Next = head
			dummyRight = dummyRight.Next
		}
		head = head.Next
	}
	dummyRight.Next = nil // 一定要断开，不然最后的链表很大可能是循环链表
	dummyLeft.Next = right.Next
	return left.Next
}
