package main

// nums 连续元素分组，小于 target 分为一组，nums[i] < target
func combine(nums []int, target int) [][]int {
	if len(nums) == 0 {
		return make([][]int, 0)
	}
	var (
		res        [][]int
		indexPoint int
		numsLength = len(nums)
	)
	for indexPoint < numsLength {
		var (
			count    int
			subArray []int
		)
		for i := indexPoint; i < numsLength; i++ {
			count += nums[i]
			if count >= target {
				indexPoint = i
				break
			}
			// 处理临界值
			if i == numsLength-1 {
				indexPoint = numsLength
			}
			subArray = append(subArray, nums[i])
		}
		res = append(res, subArray)
	}
	return res
}
