package everyday

import "math"

// 162. 寻找峰值
// https://leetcode.cn/problems/find-peak-element/description

func findPeakElement(nums []int) int {
	nl := len(nums)
	l, r := 0, nl-1
	for l < r {
		mid := l + (r-l)
		lSmaller, rSmaller := 
		nums[mid] > getIndexValue(mid-1, nl, nums),
		nums[mid] > getIndexValue(mid+1, nl, nums)
		if lSmaller && rSmaller {
			return mid
		}
		if !lSmaller {
			r = mid-1
			continue
		}
		l = mid+1	
	} 
	return -1
}

func getIndexValue(index, length int, nums []int) int {
	if index < 0 || index >= length {
		return math.MinInt
	}
	return nums[index]
}