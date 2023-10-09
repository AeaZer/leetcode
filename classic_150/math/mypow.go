package math

import "math"

// 50. Pow(x, n)
// https://leetcode.cn/problems/powx-n/?envType=study-plan-v2&envId=top-interview-150

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
