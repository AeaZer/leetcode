package main

import "sort"

func permuteUnique(nums []int) [][]int {
	length := len(nums)
	// 对数组排序
	sort.Ints(nums)
	res := make([][]int, 0)
	var temp []int
	visited := make([]bool, length)
	backTrackUnique(length, nums, temp, visited, &res)

	return res
}

func backTrackUnique(length int, nums, temp []int, visited []bool, res *[][]int) {
	if length == len(temp) {
		*res = append(*res, append([]int(nil), temp...))
		return
	}

	for index, num := range nums {
		if visited[index] || index > 0 && !visited[index-1] && num == nums[index-1] {
			continue
		}

		visited[index] = true
		temp = append(temp, num)
		backTrackUnique(length, nums, temp, visited, res)
		visited[index] = false
		if len(temp) > 1 {
			temp = temp[:len(temp)-1]
		} else {
			temp = []int{}
		}
	}

	return
}
