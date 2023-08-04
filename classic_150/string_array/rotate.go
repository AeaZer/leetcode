package string_array

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

// https://leetcode.cn/problems/rotate-array/

func rotate(nums []int, k int) {
	// 备份
	length := len(nums)
	tmp := make([]int, length)
	copy(tmp, nums)

	for i := 0; i < length; i++ {
		nums[(i+k)%length] = tmp[i]
	}
}
