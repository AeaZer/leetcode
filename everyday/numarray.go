package everyday

// 307. 区域和检索 - 数组可修改
// https://leetcode.cn/problems/range-sum-query-mutable/?envType=daily-question&envId=2023-11-14

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

type NumArray struct {
	nums []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	na := NumArray{
		nums: make([]int, n+1),
		n:    n,
	}
	for i, x := range nums {
		na.Update(i, x)
	}
	return na
}

func (n *NumArray) Update(index int, val int) {
	diff := val - (n.preSum(index+1) - n.preSum(index))
	for i := index + 1; i <= n.n; i += i & (-i) {
		n.nums[i] += diff
	}
	return

}

func (n *NumArray) SumRange(left int, right int) int {
	return n.preSum(right+1) - n.preSum(left)
}

func (n *NumArray) preSum(index int) (ans int) {
	for i := index; i > 0; i &= i - 1 {
		ans += n.nums[i]
	}
	return
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
