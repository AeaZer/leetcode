package everyday

import "sort"

// 2530. 执行 K 次操作后的最大分数
// https://leetcode.cn/problems/maximal-score-after-applying-k-operations/?envType=daily-question&envId=2023-10-18

/*给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。你的 起始分数 为 0 。

在一步 操作 中：

选出一个满足 0 <= i < nums.length 的下标 i ，
将你的 分数 增加 nums[i] ，并且
将 nums[i] 替换为 ceil(nums[i] / 3) 。
返回在 恰好 执行 k 次操作后，你可能获得的最大分数。

向上取整函数 ceil(val) 的结果是大于或等于 val 的最小整数。*/

func maxKelements(nums []int, k int) int64 {
	nl := len(nums)
	sort.Ints(nums)
	var score int64
	for k > 0 {
		score += int64(nums[nl-1])
		var smallerIndex int
		for i := nl - 2; i > 0; i-- {
			if nums[i] > nums[nl-1] {
				smallerIndex = i
			}
		}
		if smallerIndex == -1 {

		}

		k--
	}
	return 0
}
