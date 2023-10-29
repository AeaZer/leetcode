package competitions

import "sort"

// 100106. 元素和最小的山形三元组 I
// https://leetcode.cn/contest/weekly-contest-368/problems/minimum-sum-of-mountain-triplets-i/

/*给你一个下标从 0 开始的整数数组 nums 。

如果下标三元组 (i, j, k) 满足下述全部条件，则认为它是一个 山形三元组 ：

i < j < k
nums[i] < nums[j] 且 nums[k] < nums[j]
请你找出 nums 中 元素和最小 的山形三元组，并返回其 元素和 。如果不存在满足条件的三元组，返回 -1 。*/

func minimumSum(nums []int) int {
	nl := len(nums)
	if nl < 3 {
		return -1
	}

	leftDP, rightDP := make([]int, nl-1), make([]int, nl-1)
	leftDP[0], rightDP[0] = nums[0], nums[nl-1]
	for i, j := 1, nl-2; i < nl-1; {
		if nums[i] < leftDP[i-1] {
			leftDP[i] = nums[i]
		} else {
			leftDP[i] = leftDP[i-1]
		}
		if nums[j] < rightDP[i-1] {
			rightDP[i] = nums[j]
		} else {
			rightDP[i] = rightDP[i-1]
		}
		i++
		j--
	}

	sums := make([]int, 0, nl-2)
	for i := 1; i < nl-1; i++ {
		if nums[i] > leftDP[i-1] && nums[i] > rightDP[nl-i-2] {
			sums = append(sums, nums[i]+leftDP[i-1]+rightDP[nl-i-2])
		}
	}
	if len(sums) == 0 {
		return -1
	}

	sort.Ints(sums)
	return sums[0]
}
