package main

import (
	"sort"
)

func countPairs(n int, edges [][]int) int64 {
	for i := range edges {
		if edges[i][0] > edges[i][1] {
			edges[i][0], edges[i][1] = edges[i][1], edges[i][0]
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][0] != edges[j][0] {
			return edges[i][0] < edges[j][0]
		}
		return edges[i][1] < edges[j][1]
	})
	storeMap := make(map[int][]int, 0)
	visited := make(map[int]bool, n)
	for i := 0; i < len(edges); i++ {
		parentValue, childValue := edges[i][0], edges[i][1]
		visited[parentValue] = true
		visited[childValue] = true
		var flag bool
		for key, children := range storeMap {
			if key == parentValue {
				storeMap[key] = mergeSlice(storeMap[key], childValue)
				flag = true
				break
			}

			for index := range children {
				if children[index] == parentValue {
					storeMap[key] = mergeSlice(storeMap[key], childValue)
					flag = true
					break
				}
				if children[index] == childValue {
					storeMap[key] = mergeSlice(storeMap[key], parentValue)
					flag = true
					break
				}
			}
		}
		if !flag {
			storeMap[parentValue] = []int{childValue}
		}
	}

	auxiliary := make([]int64, 0)
	for i := 0; i < n; i++ {
		if !visited[i] {
			auxiliary = append(auxiliary, 1)
		}
	}

	for _, values := range storeMap {
		auxiliary = append(auxiliary, int64(len(values))+1)
	}

	var ans int64
	al := len(auxiliary)
	if al == 1 {
		return 0
	}
	for i := 0; i < al; i++ {
		for j := i + 1; j < al; j++ {
			ans += auxiliary[i] * auxiliary[j]
		}
	}

	return ans
}

func mergeSlice(nums []int, num int) []int {
	for i := range nums {
		if nums[i] == num {
			return nums
		}
	}

	return append(nums, num)
}
