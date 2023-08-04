package string_array

/*给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

0 <= j <= nums[i]
i + j < n

返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。*/

// https://leetcode.cn/problems/jump-game-ii/

// 思路：挨个跳，动态规划的思想更新每个节点最小的 step
func jump(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}

	// 好像这个 dp 是可以不需要的，可以优化删除
	dp := make([]int, length)
	for i := range nums {
		if nums[i]+i >= length-1 {
			return dp[i] + 1
		}
		step := dp[i] // 记录之前跳到当前下标的最小 step
		for j := i + 1; j < i+1+nums[i]; j++ {
			if dp[j] != 0 { // 说明之前的节点就可以跳过来，所以不用作任何操作
				continue
			}
			dp[j] = step + 1 // 跳不过来，更新跳到目标节点需要的步数
		}
	}
	return -1
}
