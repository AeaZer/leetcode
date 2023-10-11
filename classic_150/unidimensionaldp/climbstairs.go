package unidimensionaldp

// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/description/?envType=study-plan-v2&envId=top-interview-150
/*输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶*/

/*输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶*/

/*输入：n = 4
输出：5
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 1 阶 +
3. 2 阶 + 1 阶*/

func climbStairs(n int) int {
	beforeBefore, before, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		beforeBefore = before
		before = r
		r = beforeBefore + before
	}
	return r
}

func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}
