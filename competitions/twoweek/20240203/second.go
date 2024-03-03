package _0240203

import "math"

func maximumSubarraySum(nums []int, k int) int64 {
	var ans *int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {

			if abs(nums[i], nums[j]) == k {
				cs := sum(nums, i, j)
				if ans == nil {
					ans = newInt(cs)
					continue
				}
				if cs > *ans {
					ans = newInt(cs)
				}
			}

		}
	}
	if ans == nil {
		return 0
	}
	return int64(*ans)
}

func newInt(a int) *int {
	return &a
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func sum(nums []int, start, end int) int {
	var ans int
	for i := start; i <= end; i++ {
		ans += nums[i]
	}
	return ans
}

func maximumSubarraySum2(nums []int, k int) int64 {
	count := make(map[int]int)
	prefixSum := 0
	maxSum := math.MinInt

	for _, num := range nums {
		prefixSum += num
		if prefixSum == k {
			maxSum = max(maxSum, prefixSum)
		}

		if val, ok := count[prefixSum-k]; ok {
			maxSum = max(maxSum, prefixSum-val)
		}

		count[prefixSum] = max(count[prefixSum], prefixSum)
	}
	if maxSum == math.MinInt {
		return 0
	}

	return int64(maxSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
