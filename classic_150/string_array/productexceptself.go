package string_array

/*给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请不要使用除法，且在 O(n) 时间复杂度内完成此题。*/

// 1,2,3,4

// https://leetcode.cn/problems/product-of-array-except-self/

func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length <= 0 {
		return []int{}
	}

	l := make([]int, length)
	r := make([]int, length)
	l[0] = 1
	for i := 1; i < length; i++ { // [1,1*1,1*2,1*2*3]
		l[i] = nums[i-1] * l[i-1]
	}
	r[length-1] = 1
	for j := length - 2; j >= 0; j-- { // [2*3*4, 3*4, 4， 1]
		r[j] = nums[j+1] * r[j+1]
	}

	answer := make([]int, length)
	for i := 0; i < length; i++ {
		answer[i] = l[i] * r[i]
	}

	return answer
}
