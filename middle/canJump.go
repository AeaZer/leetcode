package main

func canJump(nums []int) bool {
	reach := 0
	for index, num := range nums {
		if index > reach {
			return false
		}
		reach = max(num+index, reach)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
