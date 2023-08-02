package taff

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel    = 4   // 跳表的最大层数
	probability = 0.5 // 向上提升层次的概率
)

// 跳表节点
type skipListNode struct {
	value int             // 节点值
	next  []*skipListNode // 每一层的下一个节点指针
}

// 跳表结构
type skipList struct {
	head  *skipListNode // 头节点
	level int           // 当前层数
}

// 创建新节点
func newSkipListNode(value int, level int) *skipListNode {
	return &skipListNode{
		value: value,
		next:  make([]*skipListNode, level),
	}
}

// 创建跳表
func newSkipList() *skipList {
	head := newSkipListNode(0, maxLevel)
	return &skipList{
		head:  head,
		level: 1,
	}
}

// 插入节点
func (sl *skipList) insert(value int) {
	update := make([]*skipListNode, maxLevel) // 保存需要更新的节点
	current := sl.head

	// 从最高层往下查找每一层的插入位置
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	// 生成一个随机层数
	level := randomLevel()

	// 如果生成的随机层数大于当前跳表层数，则更新当前跳表层数
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	// 创建新节点
	node := newSkipListNode(value, level)

	// 将新节点插入到每一层的链表中
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}
}

// 删除节点
func (sl *skipList) delete(value int) {
	update := make([]*skipListNode, maxLevel) // 保存需要更新的节点
	current := sl.head

	// 从最高层往下查找待删除节点的位置
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	// 获取待删除节点
	node := current.next[0]

	// 如果待删除节点存在且值与目标值相等，则进行删除操作
	if node != nil && node.value == value {
		// 遍历每一层，更新节点指针
		for i := 0; i < sl.level; i++ {
			if update[i].next[i] != node {
				break
			}
			update[i].next[i] = node.next[i]
		}

		// 更新跳表的层数
		for sl.level > 1 && sl.head.next[sl.level-
			1] == nil {
			sl.level--
		}
	}
}

// 查找节点
func (sl *skipList) search(value int) bool {
	current := sl.head

	// 从最高层往下查找目标节点
	for i := sl.level - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
	}

	// 检查目标节点是否存在
	if current.next[0] != nil && current.next[0].value == value {
		return true
	}

	return false
}

// 打印跳表
func (sl *skipList) print() {
	for i := sl.level - 1; i >= 0; i-- {
		current := sl.head.next[i]
		fmt.Printf("Level %d: ", i)
		for current != nil {
			fmt.Printf("%d ", current.value)
			current = current.next[i]
		}
		fmt.Println()
	}
}

// 生成随机层数
func randomLevel() int {
	level := 1
	for rand.Float64() < probability && level < maxLevel {
		level++
	}
	return level
}
