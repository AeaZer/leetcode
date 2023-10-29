package main

// 1726. 同积元组
// https://leetcode.cn/problems/tuple-with-same-product/description/?envType=daily-question&envId=2023-10-19

/*给你一个由 不同 正整数组成的数组 nums ，
请你返回满足 a * b = c * d 的元组 (a, b, c, d) 的数量。其中 a、b、c 和 d 都是 nums 中的元素，且 a != b != c != d 。*/

func factorial2(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial2(n-1)
}

func countCombinations(n int) int {
	return factorial2(n) / (factorial2(2) * factorial2(n-2))
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
