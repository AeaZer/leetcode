package hash

import "math"

/*给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个
不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。*/

// https://leetcode.cn/problems/contains-duplicate-ii/

// 思路：使用 map 存储 num 中的元素，第一次遍历填充 map，第二次遍历查找 map[nums[i]] 是否存在
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i := range nums {
		// 判断当前元素是否在 m 中存在并且当前的元素与 map 中相同的元素下标差是小于等于 k 的
		if v, ok := m[nums[i]]; abs(v, i) <= k && ok {
			return true
		}
		m[nums[i]] = i
	}
	return false
}

func abs(i, j int) int {
	if i == j {
		return math.MaxInt
	}
	if i > j {
		return i - j
	}
	return j - i
}
