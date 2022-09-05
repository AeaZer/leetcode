package main

func searchRange(nums []int, target int) []int {
	res := []int{
		-1, -1,
	}
	for start, end := 0, len(nums)-1; start < len(nums) && end >= 0; {
		var (
			flagStart bool
			flagEnd   bool
		)

		if nums[start] == target {
			flagStart = true
			res[0] = start
			if res[1] != -1 {
				break
			}
		}
		if nums[start] != target && !flagStart {
			start++
		}

		if nums[end] == target {
			flagEnd = true
			res[1] = end
			if res[0] != -1 {
				break
			}
		}
		if nums[end] != target && !flagEnd {
			end--
		}

	}

	return res
}
