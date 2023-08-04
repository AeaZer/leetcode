package string_array

/*给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。*/

// https://leetcode.cn/problems/jump-game/

func canJump(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}

	var maxIndex int
	for i := 0; i < len(arr); i++ {
		if i > maxIndex {
			return false
		}
		// 第 i 个元素可以挑的最大位置和之前元素可以跳的最大位置做比较
		maxIndex = max(maxIndex, arr[i]+i)
		if maxIndex >= len(arr)-1 {
			return true
		}
	}
	return false
}

func max(i, j int) int {
	if j > i {
		return j
	}
	return i
}

// more efficient
func canJump2(nums []int) bool {
	reach := 0
	for index, num := range nums {
		if index > reach {
			return false
		}
		reach = max(num+index, reach)
	}
	return true
}
