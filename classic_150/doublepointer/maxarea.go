package doublepointer

/*给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。*/

// https://leetcode.cn/problems/container-with-most-water/

// 经典双指针问题，容器装多少税取决于短板，改变的指针总是短板的指针，这样下来的结果最大值总是包含在这个思路里
func maxArea(height []int) int {
	lp, rp := 0, len(height)-1
	var maxAreas int
	for lp < rp {
		area := min(height[lp], height[rp]) * (rp - lp)
		if area > maxAreas {
			maxAreas = area
		}
		if height[lp] < height[rp] {
			lp++
			continue
		}
		rp--
	}
	return maxAreas
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
