package competitions

//给你一个下标从 0 开始、长度为 n 的整数数组 nums ，以及整数 indexDifference 和整数 valueDifference 。
//
//你的任务是从范围 [0, n - 1] 内找出  2 个满足下述所有条件的下标 i 和 j ：
//
//abs(i - j) >= indexDifference 且
//abs(nums[i] - nums[j]) >= valueDifference
//返回整数数组 answer。如果存在满足题目要求的两个下标，则 answer = [i, j] ；否则，answer = [-1, -1] 。
//如果存在多组可供选择的下标对，只需要返回其中任意一组即可。
//
//注意：i 和 j 可能 相等 。

//1 <= n == nums.length <= 100
//0 <= nums[i] <= 50
//0 <= indexDifference <= 100
//0 <= valueDifference <= 50

//1 <= n == nums.length <= 105
//0 <= nums[i] <= 109
//0 <= indexDifference <= 105
//0 <= valueDifference <= 109

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	nl := len(nums)
	for left := 0; left < nl; left++ {
		right := left + indexDifference
		for right < nl {
			if abs(nums[left], nums[right]) >= valueDifference {
				return []int{left, right}
			}
			right++
		}
	}
	return []int{-1, -1}
}

func findIndices2(nums []int, indexDifference int, valueDifference int) []int {
	nl := len(nums)
	for left := 0; left < nl-indexDifference; left++ {
		right := left + indexDifference
		for right < nl {
			if abs(nums[left], nums[right]) >= valueDifference {
				return []int{left, right}
			}
			right++
		}
	}
	return []int{-1, -1}
}
