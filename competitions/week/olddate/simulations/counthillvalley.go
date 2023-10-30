package simulations

// 2210. 统计数组中峰和谷的数量
// https://leetcode.cn/contest/weekly-contest-285/problems/count-hills-and-valleys-in-an-array/

// 给你一个下标从 0 开始的整数数组 nums 。如果两侧距 i 最近的不相等邻居的值均小于 nums[i] ，
// 则下标 i 是 nums 中，某个峰的一部分。类似地，如果两侧距 i 最近的不相等邻居的值均大于 nums[i] ，
// 则下标 i 是 nums 中某个谷的一部分。对于相邻下标 i 和 j ，如果 nums[i] == nums[j] ，
// 则认为这两下标属于 同一个 峰或谷。
//
// 注意，要使某个下标所做峰或谷的一部分，那么它左右两侧必须 都 存在不相等邻居。
//
// 返回 nums 中峰和谷的数量。

/*输入：nums = [2,4,1,1,6,5]
输出：3
解释：
在下标 0 ：由于 2 的左侧不存在不相等邻居，所以下标 0 既不是峰也不是谷。
在下标 1 ：4 的最近不相等邻居是 2 和 1 。由于 4 > 2 且 4 > 1 ，下标 1 是一个峰。
在下标 2 ：1 的最近不相等邻居是 4 和 6 。由于 1 < 4 且 1 < 6 ，下标 2 是一个谷。
在下标 3 ：1 的最近不相等邻居是 4 和 6 。由于 1 < 4 且 1 < 6 ，下标 3 符合谷的定义，但需要注意它和下标 2 是同一个谷的一部分。
在下标 4 ：6 的最近不相等邻居是 1 和 5 。由于 6 > 1 且 6 > 5 ，下标 4 是一个峰。
在下标 5 ：由于 5 的右侧不存在不相等邻居，所以下标 5 既不是峰也不是谷。
共有 3 个峰和谷，所以返回 3 。*/

func countHillValley(nums []int) int {
	numSet := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			numSet = append(numSet, nums[i])
		}
	}
	numSet = append(numSet, nums[len(nums)-1])
	nl := len(numSet)
	if nl < 3 {
		return 0
	}
	var ans int
	left, middle, right := 0, 1, 2
	for right < nl {
		if (numSet[middle] > numSet[left] && numSet[middle] > numSet[right]) ||
			(numSet[middle] < numSet[left] && numSet[middle] < numSet[right]) {
			ans++
		}
		left++
		middle++
		right++
	}
	return ans
}
