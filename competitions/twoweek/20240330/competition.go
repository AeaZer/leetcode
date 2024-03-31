package _0240330

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	ans := math.MaxInt
	for i := 0; i < len(nums); i++ {
		var orValue int
		for j := i; j < len(nums); j++ {
			orValue |= nums[j]
			if orValue >= k {
				ans = min(j-i+1, ans)
			}
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

func minimumLevels2(possible []int) int {
	dp := make([]int, len(possible))
	dp[0] = possibleScore(possible[0])
	storeSmallIndexMap := make(map[int]int)
	storeSmallIndexMap[dp[0]] = 1
	for i := 1; i < len(possible); i++ {
		dp[i] = dp[i-1] + possibleScore(possible[i])
		if _, ok := storeSmallIndexMap[dp[i]]; !ok {
			storeSmallIndexMap[dp[i]] = i + 1
		}
	}
	score := dp[len(possible)-1]/2 + 1
	if (score == 0 || score == 2) && len(possible) == 2 {
		return -1
	}
	if score <= 0 {
		return 1
	}
	ans, ok := storeSmallIndexMap[score]
	if !ok {
		return -1
	}
	return ans
}

func possibleScore(value int) int {
	if value == 0 {
		return -1
	}
	return 1
}

func minimumLevels(possible []int) int {
	length := len(possible)
	ldp := make([]int, length)
	ldp[0] = possibleScore(possible[0])

	rdp := make([]int, length)
	rdp[0] = possibleScore(possible[length-1])

	for i := 1; i < length; i++ {
		ldp[i] = ldp[i-1] + possibleScore(possible[i])

		rIndex := length - 1 - i
		rdp[i] = rdp[i-1] + possibleScore(possible[rIndex])
	}

	for i := 0; i < length-1; i++ {
		if ldp[i] > rdp[length-i-2] {
			return i + 1
		}
	}
	return -1
}

func minimumSubarrayLength2(nums []int, k int) int {
	var (
		maxIndex int
		maxValue = math.MinInt
	)
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxIndex = i
			maxValue = nums[i]
		}
	}
	if maxValue >= k {
		return 1
	}
	i, j := maxIndex-1, maxIndex+1
	for maxValue < k {
		if i >= 0 && j <= len(nums)-1 {
			if nums[i] > nums[j] {
				maxValue |= nums[i]
				i--
				continue
			}
			if nums[i] <= nums[j] {
				maxValue |= nums[j]
				j++
				continue
			}
		}
		if i >= 0 {
			maxValue |= nums[i]
			i--
			continue
		}
		if j <= len(nums)-1 {
			maxValue |= nums[j]
			j++
			continue
		}
		if i == -1 && j == len(nums) {
			return -1
		}
	}
	return j - i - 2 + 1
}
