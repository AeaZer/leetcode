package _0231028

/*给你一个下标从 0 开始的整数数组 nums 和一个整数 target 。

返回和为 target 的 nums 子序列中，子序列 长度的最大值 。如果不存在和为 target 的子序列，返回 -1 。

子序列 指的是从原数组中删除一些或者不删除任何元素后，剩余元素保持原来的顺序构成的数组。*/

func lengthOfLongestSubsequence(nums []int, target int) int {
	l := len(nums)
	var dfs func([]int, int, int)
	satisfies := make([][]int, 0)
	dfs = func(elements []int, index int, sum int) {
		if index >= l {
			return
		}
		if nums[index] == sum {
			satisfies = append(satisfies, append(elements, nums[index]))
			return
		}
		if nums[index] > sum {
			dfs(elements, index+1, sum)
		}
		dfs(append(elements, nums[index]), index+1, sum-nums[index])
		if len(elements) > 0 {
			dfs(elements[:len(elements)-1], index+1, sum-elements[len(elements)-1])
		}
	}
	for i := 0; i < l; i++ {
		dfs([]int{nums[i]}, i+1, target-nums[i])
	}
	satisfiesLen := len(satisfies)
	if satisfiesLen == 0 {
		return -1
	}
	maxLen := len(satisfies[0])
	for i := 0; i < satisfiesLen; i++ {
		if maxLen < len(satisfies[i]) {
			maxLen = len(satisfies[i])
		}
	}
	return maxLen
}
