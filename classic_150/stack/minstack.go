package stack

// NOTICE: 这道题明显使用切片更好，但是使用数组太无聊了，用链表实现一下

import "math"

type node struct {
	val      int
	nextNode *node
}

func (n node) minValue() int {
	minValue := n.val
	for n.nextNode != nil {
		if minValue > n.nextNode.val {
			minValue = n.nextNode.val
		}
		n = *n.nextNode
	}
	return minValue
}

type MinStack struct {
	node     *node
	minValue int
}

func Constructor() MinStack {
	return MinStack{
		minValue: math.MaxInt32,
	}
}

func (s *MinStack) Push(val int) {
	if val < s.minValue {
		s.minValue = val
	}
	s.node = &node{
		val:      val,
		nextNode: s.node,
	}
	s.node.minValue()
}

// Pop 删除栈顶元素
func (s *MinStack) Pop() {
	if s.node == nil {
		panic("empty stack should not pop from stack")
	}
	s.node = s.node.nextNode
	if s.node != nil {
		s.minValue = s.node.minValue()
		return
	}
	s.minValue = math.MaxInt
}

// Top 获取栈顶元素
func (s *MinStack) Top() int {
	if s.node == nil {
		panic("empty stack, cannot get top element")
	}
	return s.node.val
}

func (s *MinStack) GetMin() int {
	return s.minValue
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
