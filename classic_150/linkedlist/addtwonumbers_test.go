package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newListNode(values []int) *ListNode {
	l := new(ListNode)
	head := l
	for _, value := range values {
		head.Next = &ListNode{
			Val: value,
		}
		head = head.Next
	}
	return l.Next
}

func rangeListNode(l ListNode) []int {
	res := make([]int, 0)
	tmp := &l
	for tmp != nil {
		res = append(res, tmp.Val)
		tmp = tmp.Next
	}
	return res
}

func TestAddTwoNumbers(t *testing.T) {
	assert := assert.New(t)

	l1 := newListNode([]int{1, 2, 4})
	l2 := newListNode([]int{5, 6, 4})
	listNode := addTwoNumbers(l1, l2)
	es := rangeListNode(*listNode)
	assert.Equal([]int{6, 8, 8}, es)

	l1 = newListNode([]int{0})
	l2 = newListNode([]int{})
	listNode = addTwoNumbers(l1, l2)
	es = rangeListNode(*listNode)
	assert.Equal([]int{0}, es)

	l1 = newListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = newListNode([]int{9, 9, 9, 9})
	listNode = addTwoNumbers(l1, l2)
	es = rangeListNode(*listNode)
	assert.Equal([]int{8, 9, 9, 9, 0, 0, 0, 1}, es)
}
