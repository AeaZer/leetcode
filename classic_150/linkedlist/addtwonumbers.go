package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/*每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// https://leetcode.cn/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	tmp := head   // 临时链表
	var carry int // 携带 pre node 的相加进位数
	// 遍历链表
	for l1 != nil || l2 != nil {
		l1Num, l2Num := 0, 0
		if l1 != nil {
			l1Num = l1.Val
		}
		if l2 != nil {
			l2Num = l2.Val
		}

		sum := l1Num + l2Num + carry
		currentVal := sum % 10
		tmp.Next = &ListNode{
			Val: currentVal,
		}
		tmp = tmp.Next
		if sum > 0 {
			carry = sum / 10
		} else {
			carry = 0
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}
	// 解决最后一位先加大于等于 10 的情况，
	// 如果这个 carry 不等于 0 那么必定等于 1，因为 0 <= val <= 9
	if carry != 0 {
		tmp.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}
