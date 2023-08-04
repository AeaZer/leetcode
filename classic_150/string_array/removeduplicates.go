package string_array

/*给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。*/

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	var (
		currentAppearCount = 1 // 当前元素出现的次数
		deleteCount        int // 已删除的 num 个数
	)
	for i := 1; i < len(nums)-deleteCount; i++ {
		if nums[i] != nums[i-1] {
			currentAppearCount = 1
			continue
		}
		if currentAppearCount < 2 {
			currentAppearCount++
			continue
		}
		if currentAppearCount == 2 {
			copy(nums[i:], nums[i+1:]) // 左移，不会有 out index 的问题
			i--                        // 左移后继续判断当前下标
			deleteCount++
		}
	}
	return len(nums) - deleteCount
}
