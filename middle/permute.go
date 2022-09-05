package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	length := len(nums)
	var temp []int
	visited := make([]bool, len(nums))
	backTrack(length, nums, temp, visited, &res)

	return res
}

func backTrack(length int, nums, temp []int, visited []bool, res *[][]int) {
	if length == len(temp) {
		*res = append(*res, append([]int(nil), temp...))
		fmt.Println(temp)
		return
	}

	for index, num := range nums {
		if visited[index] {
			continue
		}

		visited[index] = true
		temp = append(temp, num)
		backTrack(length, nums, temp, visited, res)
		visited[index] = false
		if len(temp) > 1 {
			temp = temp[:len(temp)-1]
		} else {
			temp = []int{}
		}
	}

	return
}

func permute2(nums []int) [][]int {
	var res [][]int
	length := len(nums)
	var temp []int
	visited := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if length == len(temp) {
			res = append(res, append([]int(nil), temp...))
			fmt.Println(temp)
			return
		}

		for index, num := range nums {
			if visited[index] {
				continue
			}

			visited[index] = true
			temp = append(temp, num)
			dfs()
			visited[index] = false
			if len(temp) > 1 {
				temp = temp[:len(temp)-1]
			} else {
				temp = []int{}
			}
		}
	}
	dfs()

	return res
}
