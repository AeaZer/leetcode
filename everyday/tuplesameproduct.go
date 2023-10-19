package everyday

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func countCombinations(n int) int {
	return factorial(n) / (factorial(2) * factorial(n-2))
}

func tupleSameProduct(nums []int) int {
	combine := make(map[int]int)
	nl := len(nums)
	for i := 0; i < nl; i++ {
		for j := i + 1; j < nl; j++ {
			combine[nums[i]*nums[j]]++
		}
	}
	var ans int
	for _, value := range combine {
		if value > 1 {
			ans += countCombinations(value)
		}
	}
	return ans * 8
}
