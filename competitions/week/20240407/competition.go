package _0240407

import (
	"strings"
)

const (
	sortTypeASC = iota
	sortTypeDESC
)

func compared(sortType int, a, b int) bool {
	switch sortType {
	case sortTypeASC:
		return a < b
	case sortTypeDESC:
		return a > b
	}
	return false
}

func longestMonotonicSubarray(nums []int) int {
	ans := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		var (
			sortType int
			length   = 1
		)
		if nums[i] < nums[i+1] {
			sortType = sortTypeASC
		} else {
			sortType = sortTypeDESC
		}
		length++
		for j := i + 2; j < len(nums); j++ {
			if compared(sortType, nums[j-1], nums[j]) {
				length++
				continue
			}
			break
		}
		ans = max(ans, length)
	}
	return ans
}

func getSmallestString(s string, k int) string {
	const total int32 = 26
	rest := int32(k)
	var sb strings.Builder
	sb.Grow(len(s))
	for _, v := range s {
		after := v - 'a'
		before := total - after
		need := min(after, before)
		if need <= rest {
			sb.WriteByte('a')
			rest -= need
			continue
		}
		sb.WriteRune(v - rest)
		rest = 0
	}
	return sb.String()
}
