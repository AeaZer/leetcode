package _0240204

func returnToBoundaryCount(nums []int) int {
	var (
		sum int
		ans int
	)
	for _, num := range nums {
		if num == 0 {
			continue
		}
		sum += num
		if sum == 0 {
			ans++
		}
	}
	return ans
}
