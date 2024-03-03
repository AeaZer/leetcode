package _0240203

import "math"

func triangleType(nums []int) string {
	a1, a2, a3 := nums[0], nums[1], nums[2]
	if a1+a2 <= a3 || math.Abs(float64(a1-a2)) >= float64(a3) {
		return "none"
	}
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	switch len(m) {
	case 1:
		return "equilateral"
	case 2:
		return "isosceles"
	default: // case 3:
		return "scalene"
	}
}
