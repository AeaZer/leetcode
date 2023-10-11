package main

import "sort"

const (
	directionR = "R"
	directionL = "L"
)

func checkCrash(nums, directions []int) {
	valueIndexesMap := make(map[int][]int, len(nums))
	for i, num := range nums {
		valueIndexesMap[num] = append(valueIndexesMap[num], i)
	}
	if len(valueIndexesMap) == len(nums) {
		return
	}
	for _, indexes := range valueIndexesMap {
		if len(indexes) > 1 {
			for _, index := range indexes {
				directions[index] = 0 - directions[index]
			}
		}
	}
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func arrAbs(arr []int) int {
	var sum int
	for i, num := range arr {
		for j := i + 1; j < len(arr); j++ {
			sum += abs(num, arr[j])
		}
	}
	return sum
}

func sumDistance(nums []int, s string, d int) int {
	nl := len(nums)
	directions := make([]int, nl)
	for i := 0; i < nl; i++ {
		if s[i:i+1] == directionR {
			directions[i] = 1
		}
		if s[i:i+1] == directionL {
			directions[i] = -1
		}
	}
	for d > 0 {
		checkCrash(nums, directions)
		for i := 0; i < nl; i++ {
			nums[i] += directions[i]
		}
		d--
	}
	return arrAbs(nums) % (1e9 + 7)
}

func sumDistance2(nums []int, s string, d int) int {
	nl := len(nums)
	directions := make([]int, nl)
	for i := 0; i < nl; i++ {
		if s[i:i+1] == directionR {
			directions[i] = 1
		}
		if s[i:i+1] == directionL {
			directions[i] = -1
		}
	}
	for i := 0; i < nl; i++ {
		nums[i] += directions[i] * d
	}
	sort.Ints(nums)
	var res int
	const mod int = 1e9 + 7
	for i := 1; i < nl; i++ {
		res += (nums[i] - nums[i-1]) * i % mod * (nl - i) % mod
		res %= mod
	}

	return res
}
