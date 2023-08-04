package string_array

/*给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。*/

// https://leetcode.cn/problems/remove-element/

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == val {
			nums[0] = 0
			return 0
		}
		return 1
	}

	// 记录有多少个 val 被删除
	var varsCount int
	for i := 0; i < len(nums)-varsCount; i++ {
		if nums[i] == val {
			varsCount++
			copy(nums[i:], nums[i+1:]) // 复制向左移动
			i--                        // 当前下标被删除所以 下一次检查还是应该从当前下标开始
		}
	}
	copy(nums, nums[:len(nums)-varsCount])
	return len(nums) - varsCount
}
