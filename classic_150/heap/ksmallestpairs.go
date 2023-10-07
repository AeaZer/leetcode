package heap

import (
	"sort"
)

type kSmallest struct {
	num1, num2 int
	sum        int
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	arr := make([]*kSmallest, 0, len(nums1)*len(nums2))
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			arr = append(arr, &kSmallest{
				num1: nums1[i],
				num2: nums2[j],
				sum:  nums1[i] + nums2[j],
			})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].sum < arr[j].sum
	})
	res := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, []int{arr[i].num1, arr[i].num2})
	}
	return res
}
