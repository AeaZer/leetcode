package everyday

// 2735. 收集巧克力
/*
给你一个长度为 n 、下标从 0 开始的整数数组 nums ，表示收集不同巧克力的成本。
每个巧克力都对应一个不同的类型，最初，位于下标 i 的巧克力就对应第 i 个类型。
在一步操作中，你可以用成本 x 执行下述行为：
同时修改所有巧克力的类型，将巧克力的类型 ith 修改为类型 ((i + 1) mod n)th。
假设你可以执行任意次操作，请返回收集所有类型巧克力所需的最小成本*/

func sum(arr []int) int {
	var c int
	for i := 0; i < len(arr); i++ {
		c += arr[i]
	}
	return c
}

func minCost(nums []int, x int) int64 {
	nl := len(nums)
	f := make([]int, nl)
	copy(f, nums)
	ans := sum(f)
	for k := 0; k < nl; k++ {
		for i := 0; i < nl; i++ {
			f[i] = min(f[i], nums[(i+k)%nl])
		}
		ans = min(ans, k*x+sum(f))
	}
	return int64(ans)
}
