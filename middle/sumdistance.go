package main

// 2731. 移动机器人
// https://leetcode.cn/problems/movement-of-robots/description/?envType=daily-question&envId=2023-10-10

/*有一些机器人分布在一条无限长的数轴上，他们初始坐标用一个下标从 0 开始的整数数组 nums 表示。当你给机器人下达命令时，它们以每秒钟一单位的速度开始移动。

给你一个字符串 s ，每个字符按顺序分别表示每个机器人移动的方向。'L' 表示机器人往左或者数轴的负方向移动，'R' 表示机器人往右或者数轴的正方向移动。

当两个机器人相撞时，它们开始沿着原本相反的方向移动。

请你返回指令重复执行 d 秒后，所有机器人之间两两距离之和。由于答案可能很大，请你将答案对 109 + 7 取余后返回。

注意：

对于坐标在 i 和 j 的两个机器人，(i,j) 和 (j,i) 视为相同的坐标对。也就是说，机器人视为无差别的。
当机器人相撞时，它们 立即改变 它们的前进方向，这个过程不消耗任何时间。
当两个机器人在同一时刻占据相同的位置时，就会相撞。

例如，如果一个机器人位于位置 0 并往右移动，另一个机器人位于位置 2 并往左移动，下一秒，它们都将占据位置 1，并改变方向。再下一秒钟后，
第一个机器人位于位置 0 并往左移动，而另一个机器人位于位置 2 并往右移动。

例如，如果一个机器人位于位置 0 并往右移动，另一个机器人位于位置 1 并往左移动，下一秒，第一个机器人位于位置 0 并往左行驶，而另一个机器人位于位置 1 并往右移动。*/

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
