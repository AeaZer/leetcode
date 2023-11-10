package linkedlist

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/description/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode
	dummyHead := head
	for dummyHead != nil {
		// 存储下一个节点，在逻辑的最后要 dummyHead 应该是下一个节点的指针
		nextNode := dummyHead.Next
		// 将当前 dummyHead 和 preNode 关联起来
		dummyHead.Next = preNode
		// 将当前的前驱节点作为后面的后继节点
		preNode = dummyHead
		// 遍历下一个节点
		dummyHead = nextNode
	}
	return preNode
}
