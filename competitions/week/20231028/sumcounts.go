package _0231028

/*给你一个下标从 0 开始的整数数组 nums 。

定义 nums 一个子数组的 不同计数 值如下：

令 nums[i..j] 表示 nums 中所有下标在 i 到 j 范围内的元素构成的子数组
（满足 0 <= i <= j < nums.length ），那么我们称子数组 nums[i..j] 中不同值的数目为 nums[i..j] 的不同计数。
请你返回 nums 中所有子数组的 不同计数 的 平方 和。

由于答案可能会很大，请你将它对 109 + 7 取余 后返回。

子数组指的是一个数组里面一段连续 非空 的元素序列。*/

/*输入：nums = [1,2,1]
输出：15
解释：六个子数组分别为：
[1]: 1 个互不相同的元素。
[2]: 1 个互不相同的元素。
[1]: 1 个互不相同的元素。
[1,2]: 2 个互不相同的元素。
[2,1]: 2 个互不相同的元素。
[1,2,1]: 2 个互不相同的元素。
所有不同计数的平方和为 12 + 12 + 12 + 22 + 22 + 22 = 15 。*/

func sumCounts(nums []int) int {
	var mod int = 1e9 + 7
	nl := len(nums)
	dp := make([][][]int, nl)
	for i := range dp {
		dp[i] = make([][]int, nl)
	}
	for i := 0; i < nl; i++ {
		dp[i][i] = []int{nums[i]}
	}
	ans := nl % mod
	for i := 0; i < nl; i++ {
		for j := i + 1; j < nl; j++ {
			afterL := len(dp[i][j-1])
			var l int
			if elementInSlice(dp[i][j-1], nums[j]) {
				l = afterL
				dp[i][j] = dp[i][j-1]
			} else {
				l = afterL + 1
				tmp := make([]int, l)
				copy(tmp, append(dp[i][j-1], nums[j]))
				dp[i][j] = tmp
			}
			ans = (ans + l*l) % mod
		}
	}
	return ans
}

func elementInSlice(slice []int, element int) bool {
	for i := range slice {
		if slice[i] == element {
			return true
		}
	}
	return false
}
