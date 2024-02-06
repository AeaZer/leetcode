package dp

func deleteAndEarn(nums []int) int {
	var maxValue int
	for _, num := range nums {
		maxValue = max(num, maxValue)
	}
	newArr := make([]int, maxValue+1)
	for _, val := range nums {
		newArr[val] += val
	}
	return rob2(newArr)
}
