package linkedlist

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
	dummyRight.Next = nil
	dummyLeft.Next = right.Next
	return left.Next
}
