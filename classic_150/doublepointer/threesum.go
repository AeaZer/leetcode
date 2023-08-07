package doublepointer

import "sort"

/*你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。*/

// https://leetcode.cn/problems/3sum/

func threeSum(nums []int) [][]int {
	nl := len(nums)
	if nl < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < nl; i++ {
		curValue := nums[i]
		if (i > 0 && curValue == nums[i-1]) || curValue > 0 {
			continue
		}
		lp, rp := i+1, nl-1
		var f bool
		for lp < rp {
			// 左边去重
			if f && nums[lp] == nums[lp-1] {
				lp++
				continue
			}
			// 右边去重
			if f && rp < nl-1 && nums[rp] == nums[rp+1] {
				rp--
				continue
			}
			sum := nums[lp] + nums[rp]
			if sum+curValue == 0 {
				res = append(res, []int{curValue, nums[lp], nums[rp]})
				f = true
				lp++ // 也可以 rp--，无所谓，emo 去重真的烦死了
				continue
			}
			if sum+curValue > 0 {
				rp--
				continue
			}
			lp++
		}
	}

	return res
}
