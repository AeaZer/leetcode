package everyday

// 2807. 在链表中插入最大公约数
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	pre, next := head, head.Next
	for next != nil {
		midNode := &ListNode{
			Val:  getMaxDivisor(pre.Val, next.Val),
			Next: next,
		}
		pre.Next = midNode

		pre = next
		next = next.Next
	}
	return head
}

func getMaxDivisor(a, b int) int {
	minElem := min(a, b)
	for !(a%minElem == 0 && b%minElem == 0) {
		minElem--
	}
	return minElem
}
