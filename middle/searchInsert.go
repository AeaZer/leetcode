package main

func searchInsert(nums []int, target int) int {
	var insertIndex int
	var flag bool
	for index, num := range nums {
		if target == num {
			return index
		}
		if target <= num && !flag {
			if index != 0 {
				insertIndex = index
				flag = true
			} else {
				flag = true
			}
		}
		if !flag && index == len(nums)-1 {
			insertIndex = len(nums)
		}
	}

	return insertIndex
}
