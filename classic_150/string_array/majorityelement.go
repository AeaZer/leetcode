package string_array

/*给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。*/

// https://leetcode.cn/problems/majority-element/

func majorityElement(nums []int) int {
	ml := len(nums) / 2
	numsMap := make(map[int]int, ml)
	for _, num := range nums {
		numsMap[num]++
		if numsMap[num] > ml {
			return num
		}
	}
	return -1
}
