package string_array

/*实现RandomizedSet 类：
RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。*/

// https://leetcode.cn/problems/insert-delete-getrandom-o1/

import (
	"math/rand"
)

type RandomizedSet struct {
	vales []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		vales: make([]int, 0),
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if len(r.vales) == 0 {
		r.vales = []int{val}
		return true
	}
	for i := range r.vales {
		if r.vales[i] == val {
			return false
		}
	}
	// 比最小的小
	if r.vales[0] > val {
		r.vales = append([]int{val}, r.vales...)
		return true
	}
	// 比最大的大
	if r.vales[len(r.vales)-1] < val {
		r.vales = append(r.vales, val)
		return true
	}

	r.vales = append(r.vales, 0)
	for i := 0; i < len(r.vales)-1; i++ {
		if r.vales[i] < val && r.vales[i+1] > val {
			copy(r.vales[i+1:], r.vales[i:])
			r.vales[i+1] = val
			break
		}
	}
	return true
}

func (r *RandomizedSet) Remove(val int) bool {
	if len(r.vales) == 0 {
		return false
	}
	valAtIndex := -1
	for i := range r.vales {
		if r.vales[i] == val {
			valAtIndex = i
			break
		}
	}
	if valAtIndex == -1 {
		return false
	}
	// remove
	tmp := make([]int, 0, len(r.vales)-1)
	tmp = append(r.vales[:valAtIndex], r.vales[valAtIndex+1:]...)
	r.vales = tmp
	return true
}

func (r *RandomizedSet) GetRandom() int {
	if len(r.vales) == 0 {
		return 0
	}
	index := rand.Intn(len(r.vales))
	return r.vales[index]
}
