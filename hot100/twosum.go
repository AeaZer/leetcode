package hot100

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		j := len(nums) - 1
		for j > i {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
			j--
		}
	}
	return nil
}
