package main

// 先增后减 例：[]int{1,2,3,3,4,5,4,3,2,2,1}
func maxContinueArray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, 0, len(nums))
	for index, element := range nums {
		if index == 0 {
			dp[0] = 1
		}
		if element != nums[index-1] {
			dp[index] = dp[index-1] + 1
		} else {
			dp[index] = 1
		}
	}
	return maxElem(dp[0], dp[1:]...)
}

func maxElem(index0Elem int, elems ...int) int {
	maxElem := index0Elem
	for _, elem := range elems {
		if maxElem < elem {
			maxElem = elem
		}
	}
	return maxElem
}
